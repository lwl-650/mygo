package main

import (
	"fmt"
	"mygo/router"
	"mygo/util"

	"github.com/gin-gonic/gin"
)

func main() {

	// b := model.Book{}
	// fmt.Println(b)
	// b.SetBook("zs", 18)
	// fmt.Println(b)

	// 写日志
	// b.Siy()
	// b.Siy2()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// 创建一个默认的路由引擎

	// util.Tack()

	// 启动redis
	// util.InitRedis()
	// defer util.CloseRedis()

	r := gin.Default()
	r.Use(router.Content) //全局中间件
	router.IndexRoutersInit(r)
	router.HostRouters(r)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务

	maps := make(map[string]interface{})
	maps["sss"] = "sadas"

	s, _ := util.GenerateToken(maps)
	fmt.Println(s, "----------------------")
	fmt.Println(util.ConfirmToken(s, maps), "=================")

	r.Run(":8088")

	// book := util.Book{}
	// book.SetAge(18)
	// book.SetName("zs")
	// fmt.Println(book.GetAge())
	// fmt.Println(book.GetName())
	// fmt.Println(book)

}
