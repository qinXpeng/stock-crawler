package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"stock-crawler/module"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPageAction(ctx *gin.Context) {

	baseUrl := "https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=100&order=desc&order_by=percent&exchange=CN&market=CN"
	shachan := make(chan string)
	szachan := make(chan string)
	go goReq(shachan, baseUrl+"&type=sha")
	go goReq(szachan, baseUrl+"&type=sza")

	shastr := <-shachan
	szastr := <-szachan
	close(shachan)
	close(szachan)
	var tableDesert = make([]module.ResponseStockInfo, 0)
	doRes(shastr, &tableDesert)
	doRes(szastr, &tableDesert)

	sort.Slice(tableDesert, func(i, j int) bool {
		if tableDesert[i].Current <= tableDesert[j].Current {
			return true
		}
		return false
	})

	ctx.JSON(http.StatusOK, reqMulti(tableDesert))
}

func goReq(ch chan string, url string) {
	b, err := module.NetGet(url)
	if err != nil {
		fmt.Println(err)
		ch <- ""
	}
	ch <- string(b)
}

func ftos(a float64, num int) string {
	return strconv.FormatFloat(a, 'f', num, 64)
}

func doRes(str string, tableDesert *[]module.ResponseStockInfo) {

	if str == "" {
		fmt.Println("empty str")
		return
	}
	resBody := &module.ResData[*module.StockList]{}
	err := json.Unmarshal([]byte(str), resBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resBody == nil {
		fmt.Println("empty resBody")
	}
	for _, s := range resBody.Data.List {
		if s.Percent <= 5.0 || s.Percent >= 11.0 {
			continue
		}
		si := module.ResponseStockInfo{
			Symbol:             s.Symbol,
			Name:               s.Name,
			Current:            s.Current,
			Chg:                ftos(s.Chg, 2),
			Percent:            ftos(s.Percent, 2) + "%",
			CurrentYearPercent: ftos(s.CurrentYearPercent, 2) + "%",
			Volume:             fmt.Sprintf("%s万", ftos(s.Volume/10000, 3)),
			Amount:             fmt.Sprintf("%s亿", ftos(s.Amount/100000000, 3)),
			FloatShares:        ftos(s.FloatShares/100000000, 3) + "亿",
			TotalShares:        ftos(s.TotalShares/100000000, 3) + "亿",
			MarketCapital:      ftos(s.MarketCapital/100000000, 3) + "亿",
			FloatMarketCapital: ftos(s.FloatMarketCapital/100000000, 3) + "亿",
		}
		*tableDesert = append(*tableDesert, si)
	}
}

func reqSymbol(symbol string) string {
	url := fmt.Sprintf("https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s&begin=%d&period=day&type=before&count=-7&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance", symbol, time.Now().UnixNano()/1e6)
	b, err := module.NetGet(url)
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
	fmt.Println("resSlice:", len(resSlice))
	for _, s := range resSlice {
		if s.str != "" {
			rd := &module.ResData[*module.SymbolInfo]{}
			err := json.Unmarshal([]byte(s.str), rd)
			if err != nil {
				fmt.Println(s.index, err)
				continue
			}
			if rd != nil {
				cnt := 0
				for _, r := range rd.Data.Item {
					if r[7] >= 5.0 {
						cnt++
					}
				}
				if cnt >= 2 {
					mp[cnt-2] = append(mp[cnt-2], desert[s.index])
				}
			}
		}
	}
	defaultHeader := []module.ResponseHeaders{
		{Text: "股票代码", Value: "symbol"},
		{Text: "股票名称", Value: "name"},
		{Text: "当前价格", Value: "current"},
		{Text: "涨跌幅", Value: "percent"},
		{Text: "涨跌额", Value: "chg"},
		{Text: "年初至今", Value: "current_year_percent"},
		{Text: "成交量", Value: "volume"},
		{Text: "成交额", Value: "amount"},
		{Text: "流通股", Value: "float_shares"},
		{Text: "总股本", Value: "total_shares"},
		{Text: "总市值", Value: "market_capital"},
		{Text: "流通值", Value: "float_market_capital"},
	}
	resData := []module.PageResStockInfo{
		{
			Title:       "今日涨幅 " + fmt.Sprintf("%d", len(desert)),
			TableHeader: defaultHeader,
			TableDesert: desert,
		},
		{
			Title:       "七日2板 " + fmt.Sprintf("%d", len(mp[0])),
			TableHeader: defaultHeader,
			TableDesert: mp[0],
		},
		{
			Title:       "七日3板 " + fmt.Sprintf("%d", len(mp[1])),
			TableHeader: defaultHeader,
			TableDesert: mp[1],
		},
		{
			Title:       "七日4板 " + fmt.Sprintf("%d", len(mp[2])),
			TableHeader: defaultHeader,
			TableDesert: mp[2],
		},
		{
			Title:       "七日5板 " + fmt.Sprintf("%d", len(mp[3])),
			TableHeader: defaultHeader,
			TableDesert: mp[3],
		},
		{
			Title:       "七日6板 " + fmt.Sprintf("%d", len(mp[4])),
			TableHeader: defaultHeader,
			TableDesert: mp[4],
		},
		{
			Title:       "七日7板 " + fmt.Sprintf("%d", len(mp[5])),
			TableHeader: defaultHeader,
			TableDesert: mp[5],
		},
	}
	return resData
}
