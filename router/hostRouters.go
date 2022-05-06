package router

import (
	"mygo/controller/apis"

	"github.com/gin-gonic/gin"
)

func HostRouters(r *gin.Engine) {
	host := r.Group("/")
	{
		host.GET("/getCpu", apis.HostController{}.GetCpu)
	}
}