package main

import (
	"log"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/server"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {
	err := server.App.Start("127.0.0.1", "8080")
	if err != nil {
		log.Fatal(err)
	}
}
