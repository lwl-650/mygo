package apis

import (
	"mygo/model"
	"mygo/util"

	"github.com/gin-gonic/gin"
)

type XueController struct {
}

func (XueController) SetXue(c *gin.Context) {

	xue := model.Xue{}
	xue.Id = util.Snow()
	xue.Xue = "zs"

	util.DB.Create(&xue)

}
