package apis

import (
	"fmt"
	"mygo/model"
	"mygo/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (AdminController) SetAdmin(c *gin.Context) {
	A_name := c.PostForm("aname")
	A_pass := c.PostForm("apass")
	A_portrait := util.If(c.PostForm("aportrait") == "",
		"http://q.vuetitle.lwlsl.top/title.png", c.PostForm("aportrait"))
	A_grade, _ := strconv.Atoi(c.PostForm("agrade"))
	setAdmin := model.Admin{
		A_name:     A_name,
		A_pass:     A_pass,
		A_portrait: A_portrait,
		A_grade:    A_grade,
	}
	fmt.Println(setAdmin)
	util.DB.Create(&setAdmin)
}

func (AdminController) FindAdmin(c *gin.Context) {
	bb := []model.Admin{}
	util.DB.Find(&bb)
	util.Success(c, bb)

}

func (AdminController) UpdateAdmin(c *gin.Context) {
	// A_name := c.PostForm("aname")
	// A_pass := c.PostForm("apass")
	// A_portrait := util.If(c.PostForm("aportrait") == "",
	// "http://q.vuetitle.lwlsl.top/title.png", c.PostForm("aportrait"))
	// A_grade, _ := strconv.Atoi(c.PostForm("agrade"))
	// setAdmin := model.Admin{
	// 	A_name:     A_name,
	// 	A_pass:     A_pass,
	// 	A_portrait: A_portrait,
	// 	A_grade:    A_grade,
	// }
	// fmt.Println(setAdmin)
	setAdmin := model.Admin{}
	// util.DB.Create(&setAdmin)
	// util.DB.Model(&setAdmin).Update("a_name", "啦啦啦")
	util.DB.Model(&setAdmin).Where("a_id = ?", 4).Updates(map[string]interface{}{
		"a_name": "world", "a_pass": "18", "a_portrait": "5165", "a_grade": 2})
}
