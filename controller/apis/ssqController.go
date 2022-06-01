package apis

import (
	"fmt"
	"mygo/model"
	"mygo/util"

	"github.com/gin-gonic/gin"
)

type SsqController struct {
}

func (SsqController) GetSsq(c *gin.Context) {
	ssq := []model.Ssq{}

	util.RedisCore{}.Setredis()
	if util.DB.Find(&ssq).Error == nil {
		for index, item := range ssq {
			fmt.Println(index, item)
			// util.JSON(item)
		}

		util.Success(c, ssq)
	}
}
