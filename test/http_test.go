package test

import (
	"fmt"
	"stock-crawler/module"
	"testing"
)

func TestHttpGet(t *testing.T) {
	b, _ := module.HttpGet("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=5&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=1658549400483")

	fmt.Println(string(b))
}
