package apis

import (
	"mygo/model"
	"mygo/util"

	"github.com/gin-gonic/gin"
)

type RhttpController struct {
}

func (RhttpController) SetRhttp(c *gin.Context) {
	rhttp := model.R_http{R_name: "admin", R_method: c.Request.Method,
		R_url: c.FullPath(), R_actualurl: c.Request.URL.String(), R_status: c.Writer.Status(), R_time: util.GetTime(),
	}
	util.DB.Create(&rhttp)
}

func (RhttpController) FindHttpAll(c *gin.Context) {
	rhttpList := []model.R_http{}
	// var count int64

	if util.DB.Find(&rhttpList).Error == nil {
		// fmt.Println(util.DB.Find(&rhttpList).Count(&count))
		util.Success(c, rhttpList)
	} else {
		util.Error(c, -1, "err")
	}

}

func (RhttpController) DelHttpById(c *gin.Context) {
	if util.DB.Delete(&model.R_http{}, c.PostForm("rid")).Error == nil {
		util.Success(c, "")
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}

}
