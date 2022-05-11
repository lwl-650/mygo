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

	// getRhttp := []model.Rhttp{}
	// getAdmin := []model.Admin{}

	// util.DB.AutoMigrate(&getRhttp, &getAdmin)
	// util.DB.Preload("Rhttps").Find(&getAdmin)
	// util.DB.Preload("Admins").Find(&getRhttp)
	// b := model.B{
	// 	Bid:   5,
	// 	Bname: "66",
	// 	As: []model.A{
	// 		{
	// 			Aid:   6,
	// 			Aname: "55",
	// 		},
	// 	},
	// }
	// util.DB.Create(&b)
	// b := []model.B{}
	// a := []model.A{}
	// util.DB.Preload("As").Find(&b)
	// util.DB.Preload("Bs").Find(&a)
	// util.DB.AutoMigrate(&model.A{}, &model.B{})
	// util.DB.Preload("Rhttp").Find(&getAdmin)
	// info, _ := host.Info()
	// fmt.Println(info)

	// info, _ := cpu.Info() //总体信息
	// fmt.Println(info)

	// a, _ := cpu.Counts(true) //cpu逻辑数量
	// fmt.Println(a)           //4
	// a, _ = cpu.Counts(false) //cpu物理核心
	// fmt.Println(a)           //如果是2说明是双核超线程, 如果是4则是4核非超线程

	// info, _ := cpu.Times(false)
	// fmt.Println(info)

	// util.DB.AutoMigrate(&model.X{}, &model.Y{})
	// x := model.X{
	// 	Xname: "zs",
	// 	Y: []model.Y{
	// 		{
	// 			Yid:   2,
	// 			Yname: "LS",
	// 		},
	// 	},
	// }

	// util.DB.Create(&x)

	gy := []model.Y{}
	// https://www.jianshu.com/p/4c6594d2fa06  联查 一对一地址
	util.DB.Preload("X").Find(&gy)
	util.Success(c, gy)
}
