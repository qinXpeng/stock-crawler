package test

import (
	"encoding/json"
	"fmt"
	"stock-crawler/module"
	"testing"
)

func TestInitRegister(t *testing.T){
	module.InitRegister("../conf/dowstream.toml")
	b,_:=json.Marshal(module.GlobalRegister)
	fmt.Println(string(b))
}
