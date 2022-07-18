package main

import (
	"fmt"
	"stock-crawler/action"
	"stock-crawler/module"

	"github.com/gin-gonic/gin"
)

func initConf() {
	module.InitGlobalConf("./conf/setting.toml")
}

func main() {
	initConf()
	httpserver(module.GlobalConfig.MainConf.Port)
}

func httpserver(port string) {
	if port == "" {
		panic("empty port")
	}
	fmt.Println("listen port:", port)
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	app.GET("/stockserver/getpage", action.GetPageAction)
	app.Run(":" + port)

}
