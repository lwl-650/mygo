package apis

import (
	"fmt"
	"mygo/model"
	"mygo/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RhttpController struct {
}

func (RhttpController) SetRhttp(c *gin.Context) {
	rhttp := model.Rhttp{R_name: "admin", R_method: c.Request.Method,
		R_url: c.FullPath(), R_actualurl: c.Request.URL.String(), R_status: c.Writer.Status(), R_time: util.GetTime(),
	}
	util.DB.Create(&rhttp)
}

func (RhttpController) FindHttpAll(c *gin.Context) {
	rhttpList := []model.Rhttp{}
	fmt.Println(c.Query("current"), c.Query("pageSize"), "=====")
	var count int64
	rmap := make(map[string]interface{})
	rpage, _ := strconv.Atoi(c.Query("current"))
	rnum, _ := strconv.Atoi(c.Query("pageSize"))
	username := c.Query("username")
	method := c.Query("method")
	status := c.Query("status")

	fmt.Println(rpage, rnum, "+++++")
	rnum = util.If(rnum == 0, 10, rnum)
	getPage := util.If(rpage > 1, (rpage-1)*rnum, 0)
	fmt.Println(username, method, status)
	fmt.Println("从什么位置开始=>", getPage, "查询多少条=>", rnum)
	if util.DB.Where("r_name LIKE ? ", "%"+username+"%").Limit(rnum).Offset(getPage).Find(&rhttpList).Error == nil {
		util.DB.Where("r_name LIKE ?", "%"+username+"%").Find(&model.Rhttp{}).Count(&count)
		rmap["data"] = rhttpList
		rmap["count"] = count
		util.Success(c, rmap)
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}
}

func (RhttpController) DelHttpById(c *gin.Context) {
	if util.DB.Delete(&model.Rhttp{}, c.PostForm("rid")).Error == nil {
		util.Success(c, "")
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}

}

func (RhttpController) Getar(c *gin.Context) {
	// fmt.Println(GetCpuPercent())

	// getRhttp := model.Rhttp{}
	// getAdmin := model.Admin{}

	// util.DB.AutoMigrate(&getRhttp, &getAdmin)

	// user := []model.Rhttp{}
	// util.DB.Preload("Rhttp").Find(&getAdmin)
	util.Success(c, "")

}
