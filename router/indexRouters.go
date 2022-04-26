package router

import (
	"fmt"
	"mygo/controller/apis"

	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {
	index := r.Group("/")
	{
		index.GET("/", func(ctx *gin.Context) {
			fmt.Println(ctx.FullPath())
			fmt.Println("中间")

		}, apis.UserController{}.Index)
		index.GET("/add", func(ctx *gin.Context) {
			fmt.Println(ctx.FullPath())
			fmt.Println(ctx.Request.Method)
			fmt.Println(ctx.GetTime("123456"))
			fmt.Println("中间")

		}, apis.UserController{}.Add)
		index.POST("/find", apis.UserController{}.FindUser)
		index.POST("/findId", apis.UserController{}.FindById)
		index.POST("/addUser", apis.UserController{}.AddUser)
		index.POST("/login", func(ctx *gin.Context) {

			fmt.Println("中间")

		}, apis.UserController{}.Login)
		index.POST("/delById", apis.UserController{}.DelById)
	}

	other := r.Group("/")
	{
		other.POST("/ssq", apis.OtherController{}.Ssq)
		other.POST("/getssq", apis.OtherController{}.GetSsq)
	}
}
