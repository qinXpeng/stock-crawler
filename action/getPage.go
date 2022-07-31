package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stock-crawler/module"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPageAction(ctx *gin.Context) {
	tableDesert := getStockDataAll()
	if len(tableDesert) == 0 {
		ctx.JSON(http.StatusInternalServerError, "empty")
		return
	}
	ctx.JSON(http.StatusOK, reqMulti(tableDesert))
}

func ftos(a float64, num int) string {
	return strconv.FormatFloat(a, 'f', num, 64)
}

func stof(f float64, num int) float64 {
	v, _ := strconv.ParseFloat(ftos(f, 3), 64)
	return v
}

func reqSymbol(symbol string) string {
	baseUrl := module.GlobalRegister.Servers["Stock_Kilne"].Url
	reqUrl := fmt.Sprintf("%s?symbol=%s&begin=%d&period=day&type=before&count=-10", baseUrl, symbol, time.Now().UnixNano()/1e6)
	b, err := module.HttpGet(reqUrl)
	if err != nil {
		return ""
	}
	return string(b)
}

func reqMulti(desert []module.ResponseStockInfo) []module.PageResStockInfo {
	wg := &sync.WaitGroup{}
	type strInfo struct {
		index int
		str   string
	}
	ch := make(chan *strInfo)
	okch := make(chan bool)
	resSlice := make([]*strInfo, 0)
	godo := func() {
		for {
			select {
			case si := <-ch:
				resSlice = append(resSlice, si)
			case <-okch:
				fmt.Println("task ok")
				return
			}
		}
	}
	go godo()

	n := len(desert)
	var today []module.ResponseStockInfo
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int, sy string) {
			defer wg.Done()
			si := &strInfo{
				index: id,
				str:   reqSymbol(sy),
			}
			ch <- si
		}(i, desert[i].Symbol)
	}
	wg.Wait()
	okch <- true
	close(okch)
	close(ch)
	var mp [6][]module.ResponseStockInfo
	for _, s := range resSlice {
		if s.str != "" {
			rd := &module.ResData[*module.SymbolInfo]{}
			err := json.Unmarshal([]byte(s.str), rd)
			if err != nil {
				fmt.Println(s.index, s.str, err)
				continue
			}
			if rd != nil {
				cnt := 0
				lsi := 0
				for i, r := range rd.Data.Item {
					if r[7] >= 4.0 {
						cnt++
						if i >= 5 {
							lsi++
						}
					}
					if i == len(rd.Data.Item)-1 && (r[7] >= 8.0 || (r[7] >= 4.0 && strings.Index(desert[s.index].Name, "ST") != -1)) {
						today = append(today, desert[s.index])
					}
				}
				if cnt >= 2 {
					if cnt <= 3 && lsi != cnt && len(rd.Data.Item) >= 5{
						continue
					}
					id := cnt - 2
					if id > 5 {
						id = 5
					}
					mp[id] = append(mp[id], desert[s.index])
				}
			}
		}
	}
	defaultHeader := []module.ResponseHeaders{
		{Text: "股票代码", Value: "symbol"},
		{Text: "股票名称", Value: "name"},
		{Text: "当前价格", Value: "current"},
		{Text: "涨跌幅(%)", Value: "percent"},
		{Text: "涨跌额", Value: "chg"},
		{Text: "年初至今(%)", Value: "current_year_percent"},
		{Text: "成交量(万)", Value: "volume"},
		{Text: "成交额(亿)", Value: "amount"},
		{Text: "流通股(亿)", Value: "float_shares"},
		{Text: "总股本(亿)", Value: "total_shares"},
		{Text: "总市值(亿)", Value: "market_capital"},
		{Text: "流通值(亿)", Value: "float_market_capital"},
	}
	var resData []module.PageResStockInfo
	for i := 5; i >= 0; i-- {

		if i == 5 {
			resData = append(resData, module.PageResStockInfo{
				Title:       fmt.Sprintf("十日七板+ %d", len(mp[i])),
				TableHeader: defaultHeader,
				TableDesert: mp[i],
			})
		} else {
			if i <= 1 {
				resData = append(resData, module.PageResStockInfo{
					Title:       fmt.Sprintf("五日%d板 %d", i+2, len(mp[i])),
					TableHeader: defaultHeader,
					TableDesert: mp[i],
				})
			} else {
				resData = append(resData, module.PageResStockInfo{
					Title:       fmt.Sprintf("十日%d板 %d", i+2, len(mp[i])),
					TableHeader: defaultHeader,
					TableDesert: mp[i],
				})
			}
		}
	}
	resData = append(resData, module.PageResStockInfo{
		Title:       fmt.Sprintf("今日涨幅 %d", len(today)),
		TableHeader: defaultHeader,
		TableDesert: today,
	})
	return resData
}

func isHSstock(symbol string) bool {
	if len(symbol) < 4 {
		return false
	}
	return symbol[3] == '0' && (symbol[2] == '0' || symbol[2] == '6')
}

func getStockDataAll() []module.ResponseStockInfo {
	par_desc := "page=1&size=200&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz"
	par_asc := "page=1&size=200&order=asc&orderby=percent&order_by=percent&market=CN&type=sh_sz"
	var tableDesert []module.ResponseStockInfo
	s1, err := getHSPage(par_desc)
	if err == nil {
		tableDesert = append(tableDesert, s1...)
	}
	s2, err := getHSPage(par_asc)
	if err == nil {
		tableDesert = append(tableDesert, s2...)
	}
	return tableDesert
}

func getHSPage(params string) ([]module.ResponseStockInfo, error) {
	baseUrl := module.GlobalRegister.Servers["Stock_HS"].Url
	reqUrl := baseUrl + "?" + params
	b, err := module.HttpGet(reqUrl)
	if err != nil {
		return nil, err
	}
	resBody := &module.ResData[*module.StockList]{}
	err = json.Unmarshal(b, resBody)
	if err != nil {
		return nil, err
	}
	var tableDesert []module.ResponseStockInfo
	for _, s := range resBody.Data.List {
		if isHSstock(s.Symbol) {
			si := module.ResponseStockInfo{
				Symbol:             s.Symbol,
				Name:               s.Name,
				Current:            s.Current,
				Chg:                ftos(s.Chg, 2),
				Percent:            s.Percent,
				CurrentYearPercent: s.CurrentYearPercent,
				Volume:             stof(s.Volume/10000, 3),
				Amount:             stof(s.Amount/100000000, 3),
				FloatShares:        stof(s.FloatShares/100000000, 3),
				TotalShares:        stof(s.TotalShares/100000000, 3),
				MarketCapital:      stof(s.MarketCapital/100000000, 3),
				FloatMarketCapital: stof(s.FloatMarketCapital/100000000, 3),
			}
			tableDesert = append(tableDesert, si)
		}
	}
	return tableDesert, nil
}
