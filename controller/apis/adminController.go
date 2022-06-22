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
	if util.DB.Create(&setAdmin) == nil {
		util.Success(c, util.ApiCode.SUCCESS)
	}
}

func (AdminController) FindAdmin(c *gin.Context) {
	admin := []model.Admin{}
	a := []model.A{}
	subQuery1 := util.DB.Model(&admin).Select("a_id")
	subQuery2 := util.DB.Model(&a).Select("aid")
	util.DB.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&admin)
	// fmt.Println(subQuery1, "=======")

	// if util.DB.Find(&admin).Error == nil {
	util.Success(c, admin)
	// }

}

func (AdminController) UpdateAdmin(c *gin.Context) {
	A_name := c.PostForm("aname")
	A_pass := c.PostForm("apass")
	A_portrait := util.If(c.PostForm("aportrait") == "",
		"http://q.vuetitle.lwlsl.top/title.png", c.PostForm("aportrait"))
	A_grade, _ := strconv.Atoi(c.PostForm("agrade"))
	setAdmin := model.Admin{}
	util.DB.Model(&setAdmin).Where("a_id = ?", 4).Updates(map[string]interface{}{
		"a_name": A_name, "a_pass": A_pass, "a_portrait": A_portrait, "a_grade": A_grade})
}

func (AdminController) FindAdminByLogin(c *gin.Context) {
	admin := model.Admin{}
	setLogin := model.Loginadmin{}
	aname := c.PostForm("aname")
	apass := c.PostForm("apass")
	setLogin.Lname = aname
	setLogin.Ltime = util.GetTime()
	if aname != "" && apass != "" {
		util.DB.Where("a_name=?", aname).First(&admin)
		fmt.Println(admin.A_pass)
		if admin.A_pass == apass {
			util.Success(c, admin)
			util.DB.Create(&setLogin)
		} else {
			util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
		}
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}

}
