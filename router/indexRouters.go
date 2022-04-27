package router

import (
	"mygo/controller/apis"

	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {
	index := r.Group("/")
	{
		// index.Use(Content) //路由分组中间件
		index.GET("/", apis.UserController{}.Index)
		index.GET("/add", apis.UserController{}.Add)
		index.POST("/find", apis.UserController{}.FindUser)
		index.POST("/findId", apis.UserController{}.FindById)
		index.POST("/addUser", apis.UserController{}.AddUser)
		index.POST("/login", apis.UserController{}.Login)
		index.POST("/delById", apis.UserController{}.DelById)
	}
	rhttp := r.Group("/")
	{
		rhttp.POST("/findRhttp", apis.RhttpController{}.FindHttpAll)
		rhttp.POST("/delByRid", apis.RhttpController{}.DelHttpById)
	}
	tes := r.Group("/")
	{
		tes.POST("/getArr", apis.TesController{}.GetArr)
		tes.POST("/getMap", apis.TesController{}.GetMap)
		tes.POST("/getMapArr", apis.TesController{}.GetMapArr)
		tes.POST("/getArrMap", apis.TesController{}.GetArrMap)
	}
	other := r.Group("/")
	{
		other.GET("/ssq", apis.OtherController{}.Ssq)
		other.POST("/getssq", apis.OtherController{}.GetSsq)
	}
}
