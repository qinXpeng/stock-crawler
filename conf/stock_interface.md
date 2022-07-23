# stock 股票爬虫接口

## 雪球沪深榜单接口

```shell
method: GET
host: https://xueqiu.com/service/v5/stock/screener/quote/list
path: page=1&size=30&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=1658549400483
:type = sh_sz | sha | sza
: order = desc | asc
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
- response 详情
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

## 查看股票行业版块

```shell
method: GET
host:https://xueqiu.com/stock/industry/stockList.json
path: code=SZ300554&type=1&size=100
```

- response
```json
{
    "stockname":"",
    "platename":"机械设备",
    "industrystocks":[
        {
            "symbol":"SZ002204",
            "code":"002204",
            "name":"大连重工",
            "exchange":"SZ",
            "current":"9.53",
            "percentage":"10.05",
            "change":"0.87",
            "volume":"403817359",
            "pe_ttm":"117.98",
            "marketCapital":"1.8405956405E10"
        }
    ],
    "exchange":"CN",
    "code":"SZ300554",
    "industryname":"机械设备"
}
```

## 查看个股明细

```shell
method: GET
host:https://stock.xueqiu.com/v5/stock/quote.json
path: symbol=SZ300554&extend=detail
```
- response

```json
{
    "data":{
        "market":{
            "status_id":8,
            "region":"CN",
            "status":"休市",
            "time_zone":"Asia/Shanghai",
            "time_zone_desc":null,
            "delay_tag":0
        },
        "quote":{
            "current_ext":21.08,
            "symbol":"SZ300554",
            "volume_ext":0,
            "high52w":21.08,
            "delayed":0,
            "type":11,
            "tick_size":0.01,
            "float_shares":65334054,
            "limit_down":14.06,
            "no_profit":"N",
            "high":21.08,
            "float_market_capital":1377241858,
            "timestamp_ext":1658475294000,
            "lot_size":100,
            "lock_set":null,
            "weighted_voting_rights":"N",
            "chg":3.51,
            "eps":-0.92,
            "last_close":17.57,
            "profit_four":-86250895.83,
            "volume":12521969,
            "volume_ratio":3.59,
            "profit_forecast":-22577737,
            "turnover_rate":19.17,
            "low52w":11.7,
            "name":"三超新材",
            "exchange":"SZ",
            "pe_forecast":-87.403,
            "total_shares":93613367,
            "status":1,
            "is_vie_desc":"否",
            "security_status":null,
            "code":"300554",
            "goodwill_in_net_assets":null,
            "avg_price":20.06,
            "percent":19.98,
            "weighted_voting_rights_desc":"无差异",
            "amplitude":15.25,
            "current":21.08,
            "is_vie":"N",
            "current_year_percent":27.84,
            "issue_date":1492704000000,
            "sub_type":"3",
            "low":18.4,
            "is_registration_desc":"否",
            "no_profit_desc":"已盈利",
            "market_capital":1973369776,
            "dividend":0,
            "dividend_yield":0,
            "currency":"CNY",
            "navps":4.62,
            "profit":-75010965.42,
            "timestamp":1658475294000,
            "pe_lyr":-26.308,
            "amount":251190955.44,
            "pledge_ratio":null,
            "traded_amount_ext":0,
            "is_registration":"N",
            "pb":4.563,
            "limit_up":21.08,
            "pe_ttm":-22.879,
            "time":1658475294000,
            "open":18.54
        },
        "others":{
            "pankou_ratio":100,
            "cyb_switch":true
        },
        "tags":[

        ]
    },
    "error_code":0,
    "error_description":""
}
```

## 查看个股日线图

```shell
method: GET
host: https://stock.xueqiu.com/v5/stock/chart/kline.json
path: symbol=SZ300554&begin=1658638635370&period=day&type=before&count=-1&indicator=
```

- response

```json
{
    "data":{
        "symbol":"SZ300554",
        "column":[
            "timestamp",
            "volume",
            "open",
            "high",
            "low",
            "close",
            "chg",
            "percent",
            "turnoverrate",
            "amount",
            "volume_post",
            "amount_post",
            "pe",
            "pb",
            "ps",
            "pcf",
            "market_capital",
            "balance",
            "hold_volume_cn",
            "hold_ratio_cn",
            "net_volume_cn",
            "hold_volume_hk",
            "hold_ratio_hk",
            "net_volume_hk"
        ],
        "item":[
            [
                1658419200000,
                12521969,
                18.54,
                21.08,
                18.4,
                21.08,
                3.51,
                19.98,
                19.14,
                251190955,
                null,
                null,
                -22.8794,
                4.56,
                7.5478,
                38.7448,
                1973369776.36,
                null,
                null,
                null,
                null,
                null,
                null,
                null
            ]
        ]
    },
    "error_code":0,
    "error_description":""
}
```

## 查看板块下股票详情

```shell
method: GET
host: https://stock.xueqiu.com/v5/stock/screener/quote/list.json
path: index_symbol=BK0400&market=CN&order=desc&order_by=percent&page=1&size=5000&type=sh_sz&_=1658577566102
index_symbol: 板块代号
order_by:limitup_days 连板
```

- response
```json
{
    "data":{
        "count":146,
        "list":[
            {
                "symbol":"SZ300356",
                "net_profit_cagr":null,
                "north_net_inflow":null,
                "ps":2.743,
                "type":11,
                "percent":11.49,
                "has_follow":false,
                "tick_size":0.01,
                "pb_ttm":1.4274,
                "float_shares":391427562,
                "current":2.62,
                "amplitude":11.06,
                "pcf":24.0581,
                "current_year_percent":-22.02,
                "float_market_capital":1025540212,
                "north_net_inflow_time":null,
                "market_capital":1068710104,
                "dividend_yield":0,
                "lot_size":100,
                "roe_ttm":-25.056664174668697,
                "total_percent":-49.38,
                "percent5m":0,
                "income_cagr":1.2840337121817313,
                "amount":55888113.54,
                "chg":0.27,
                "issue_date_ts":1349712000000,
                "eps":-0.53,
                "main_net_inflows":-598207,
                "volume":22148331,
                "volume_ratio":2.18,
                "pb":1.424,
                "followers":16140,
                "turnover_rate":5.66,
                "first_percent":20.9,
                "name":"*ST光一",
                "pe_ttm":null,
                "total_shares":407904620,
                "limitup_days":0
            }
        ]
    },
    "error_code":0,
    "error_description":""
}
```

## 东方财富龙虎榜
```json
https://data.eastmoney.com/stock/tradedetail.html
```

## 东方财富营业部
```json
method: GET
https://datacenter-web.eastmoney.com/api/data/v1/get?
sortColumns=TOTAL_BUYER_SALESTIMES_1DAY,OPERATEDEPT_CODE&sortTypes=-1,1&pageSize=50&pageNumber=1&reportName=RPT_RATEDEPT_RETURNT_RANKING&columns=ALL&source=WEB&client=WEB&filter=(STATISTICSCYCLE="02")
```

- response

```json
{
    "version":"508eafc34dc57c12b560b8b3429c62a8",
    "result":{
        "pages":3239,
        "data":[
            {
                "OPERATEDEPT_CODE":"10472087",
                "OPERATEDEPT_NAME":"东方财富证券拉萨东环路第二证券营业部",
                "STATISTICSCYCLE":"02",
                "AVERAGE_INCREASE_1DAY":-1.821278552215,
                "RISE_PROBABILITY_1DAY":36.25678119349,
                "TOTAL_BUYER_SALESTIMES_1DAY":1106,
                "AVERAGE_INCREASE_2DAY":-2.782328786191,
                "RISE_PROBABILITY_2DAY":33.670653173873,
                "TOTAL_BUYER_SALESTIMES_2DAY":1087,
                "AVERAGE_INCREASE_3DAY":-3.370708821707,
                "RISE_PROBABILITY_3DAY":30.597014925373,
                "TOTAL_BUYER_SALESTIMES_3DAY":1072,
                "AVERAGE_INCREASE_5DAY":-4.036397256444,
                "RISE_PROBABILITY_5DAY":31.884057971015,
                "TOTAL_BUYER_SALESTIMES_5DAY":1035,
                "AVERAGE_INCREASE_10DAY":-5.551605445249,
                "RISE_PROBABILITY_10DAY":28.950159066808,
                "TOTAL_BUYER_SALESTIMES_10DAY":943,
                "OPERATEDEPT_CODE_OLD":"80441008"
            }
        ],
        "count":3239
    },
    "success":true,
    "message":"ok",
    "code":0
}
```

## 东方财富各营业部详情

```json
https://datacenter-web.eastmoney.com/api/data/v1/get?
sortColumns=TRADE_DATE,SECURITY_CODE&sortTypes=-1,1&pageSize=50&pageNumber=1&reportName=RPT_OPERATEDEPT_TRADE_DETAILS&columns=ALL&filter=(OPERATEDEPT_CODE="10045025")&source=WEB&client=WEB
```

- response

```json
{
    "version": "2478f3fc7fbc5f403309976c421d4d07",
    "result": {
        "pages": 6775,
        "data": [
            {
                "OPERATEDEPT_CODE": "10045025",
                "OPERATEDEPT_NAME": "中信证券股份有限公司上海溧阳路证券营业部",
                "TRADE_DATE": "2022-07-22 00:00:00",
                "D1_CLOSE_ADJCHRATE": null,
                "D2_CLOSE_ADJCHRATE": null,
                "D3_CLOSE_ADJCHRATE": null,
                "D5_CLOSE_ADJCHRATE": null,
                "D10_CLOSE_ADJCHRATE": null,
                "SECURITY_CODE": "605088",
                "SECURITY_NAME_ABBR": "冠盛股份",
                "ACT_BUY": null,
                "ACT_SELL": 23263522,
                "NET_AMT": -23263522,
                "EXPLANATION": "有价格涨跌幅限制的日换手率达到20%的前五只证券",
                "D20_CLOSE_ADJCHRATE": null,
                "D30_CLOSE_ADJCHRATE": null,
                "SECUCODE": "605088.SH",
                "OPERATEDEPT_CODE_OLD": "80032107",
                "ORG_NAME_ABBR": "中信证券上海溧阳路证券营业部",
                "CHANGE_RATE": -1.806
            }
        ],
        "count": 6775
    },
    "success": true,
    "message": "ok",
    "code": 0
}
```
