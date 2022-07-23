# stock 股票爬虫接口

## 雪球沪深榜单接口

```shell
method: GET
host: https://xueqiu.com/service/v5/stock/screener/quote/list
path: page=1&size=30&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=1658549400483
:type = sh_sz | sha | sza 
```

- response 
```json
{
    "data":{
        "count":4819,
        "list":[
            {
                "symbol":"SH688375",
                "net_profit_cagr":18.92481453595456,
                "north_net_inflow":null,
                "ps":13.9423,
                "type":82,
                "percent":35.43,
                "has_follow":false,
                "tick_size":0.01,
                "pb_ttm":7.108,
                "float_shares":30428118,
                "current":95.99,
                "amplitude":11.13,
                "pcf":26.1759,
                "current_year_percent":35.43,
                "float_market_capital":2920795047,
                "north_net_inflow_time":null,
                "market_capital":38396959900,
                "dividend_yield":null,
                "lot_size":1,
                "roe_ttm":null,
                "total_percent":2.12,
                "percent5m":0.14,
                "income_cagr":23.774544727917824,
                "amount":2047061815,
                "chg":25.11,
                "issue_date_ts":1658419200000,
                "eps":1.14,
                "main_net_inflows":0,
                "volume":21246668,
                "volume_ratio":null,
                "pb":7.11,
                "followers":1128,
                "turnover_rate":69.83,
                "first_percent":35.43,
                "name":"N国博",
                "pe_ttm":93.964,
                "total_shares":400010000,
                "limitup_days":0
            }
        ]
    },
    "error_code":0,
    "error_description":""
}
```
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

