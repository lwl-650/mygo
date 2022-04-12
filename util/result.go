package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//返回的结果：
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Msg = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

//出错
func Error(c *gin.Context, code int, msg string) {
	res := Result{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}

func Other(c *gin.Context, code int, msg string, data interface{}) {
	res := Result{}
	res.Code = code
	res.Msg = msg
	res.Data = data
	c.JSON(http.StatusOK, res)
}
