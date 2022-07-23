package module

type MainConf struct {
	Port string `toml:"port"`
}

type SetConfig struct {
	MainConf   MainConf          `toml:"main_conf"`
	HttpHeader map[string]string `toml:"http_header"`
}
type ServerConf struct {
	Url    string `toml:"url"`
	Method string `toml:"method"`
}
type RegisterConfig struct {
	Servers map[string]ServerConf `toml:"servers"`
	Headers map[string]string     `toml:"headers"`
}

type ResStockInfo struct {
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Current            float64 `json:"current"`
	Chg                float64 `json:"chg"`
	Percent            float64 `json:"percent"`
	CurrentYearPercent float64 `json:"current_year_percent"`
	Volume             float64 `json:"volume"`
	Amount             float64 `json:"amount"`
	TurnoverRate       float64 `json:"turnover_rate"`
	PeTtm              float64 `json:"pe_ttm"`
	FloatMarketCapital float64 `json:"float_market_capital"`
	MarketCapital      float64 `json:"market_capital"`
	VolumeRatio        float64 `json:"volume_ratio"`
	DividendYield      float64 `json:"dividend_yield"`
	FloatShares        float64 `json:"float_shares"`
	TotalShares        float64 `json:"total_shares"`
}

type StockList struct {
	List []ResStockInfo `json:"list"`
}

type SymbolInfo struct {
	Symbol string      `json:"symbol"`
	Column []string    `json:"column"`
	Item   [][]float64 `json:"item"`
}

type ResData[T comparable] struct {
	Data             T      `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type ResponseHeaders struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type ResponseStockInfo struct {
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Current            float64 `json:"current"`
	Chg                string  `json:"chg"`
	Percent            float64 `json:"percent"`
	CurrentYearPercent float64 `json:"current_year_percent"`
	Volume             float64 `json:"volume"`
	Amount             float64 `json:"amount"`
	TurnoverRate       float64 `json:"turnover_rate"`
	PeTtm              float64 `json:"pe_ttm"`
	FloatMarketCapital float64 `json:"float_market_capital"`
	MarketCapital      float64 `json:"market_capital"`
	VolumeRatio        float64 `json:"volume_ratio"`
	DividendYield      float64 `json:"dividend_yield"`
	FloatShares        float64 `json:"float_shares"`
	TotalShares        float64 `json:"total_shares"`
}

type PageResStockInfo struct {
	Title       string              `json:"title"`
	TableHeader []ResponseHeaders   `json:"tableheader"`
	TableDesert []ResponseStockInfo `json:"tabledesert"`
}

/*
symbol 股票代码
name 股票名称
current 当前价格
chg 涨跌额 +-
percent 涨跌幅 %
current_year_percent 年初至今 %
volume 成交量 万
amount 成交额 亿
turnover_rate换手率 %
pe_ttm 市赢率
float_market_capital 流通值 亿
market_capital 总市值 亿
volume_ratio 量比
dividend_yield 股息率 %
float_shares 流通股  亿
total_shares 总股本 亿

https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=100&order=desc&order_by=percent&exchange=CN&market=CN&type=sha&_=1658060247754

https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=100&order=desc&order_by=percent&exchange=CN&market=CN&type=sza&_=1658060275250

https://xueqiu.com/S/SZ002594


https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=SZ002090&begin=1658148203503&period=day&type=before&count=-7&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance

*/
