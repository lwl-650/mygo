package main

import (
	"mygo/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()

	router.IndexRoutersInit(r)

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8088")

	// book := util.Book{}
	// book.SetAge(18)
	// book.SetName("zs")
	// fmt.Println(book.GetAge())
	// fmt.Println(book.GetName())
	// fmt.Println(book)
}