package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stock-crawler/module"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) {
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
	baseUrl := "https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=60&order=desc&order_by=percent&exchange=CN&market=CN"
	shachan := make(chan string)
	szachan := make(chan string)
	go goReq(shachan, baseUrl+"&type=sha")
	go goReq(szachan, baseUrl+"&type=sza")

	shastr := <-shachan
	szastr := <-szachan
	close(shachan)
	close(szachan)
	var tableDesert []module.ResponseStockInfo
	doRes(shastr, &tableDesert)
	doRes(szastr, &tableDesert)
	resData := []module.PageResStockInfo{
		{
			Title:       "今日涨幅",
			TableHeader: defaultHeader,
			TableDesert: tableDesert,
		},
	}
	ctx.JSON(http.StatusOK, resData)
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
		si := module.ResponseStockInfo{
			Symbol:             s.Symbol,
			Name:               s.Name,
			Current:            ftos(s.Current, 2),
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
