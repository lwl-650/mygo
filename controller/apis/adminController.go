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
	// rhttp := []model.Rhttp{}
	// a := []model.A{}
	// subQuery1 := util.DB.Model(&admin).Select("a_id")
	// subQuery2 := util.DB.Model(&a).Select("aid")
	// subQuery2 := util.DB.Model(&rhttp).Select("r_id")
	util.DB.Raw("SELECT a.a_id, a.a_name,a.a_pass , b.r_time,b.r_id FROM admin as a ,rhttp as b WHERE a.a_id = b.r_id AND a.a_id = ?", 3).Scan(&admin)
	// util.DB.Raw("SELECT a.a_id, a.a_name,a.a_pass , b.r_time,b.r_id FROM admin as a ,rhttp as b WHERE a.a_id = ?", 1).Scan(&admin)

	// util.DB.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&admin).Where("a_id = ? ", "aid")
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
	auser := make(map[string]interface{})

	// mapss := make(map[string]interface{})
	// token, _ := util.GenerateToken(auser)

	// fmt.Println(util.ConfirmToken(s, mapss), "=================")

	admin := model.Admin{}
	setLogin := model.Loginadmin{}
	aname := c.PostForm("aname")
	apass := c.PostForm("apass")
	setLogin.Lname = aname
	setLogin.Ltime = util.GetTime()
	if aname != "" && apass != "" {
		util.DB.Where("a_name=?", aname).First(&admin)
		fmt.Println("🐱‍🏍 => file: adminController.go => line 73 => func => aname", aname)
		fmt.Println(admin.A_pass, "888888888888888888")
		if admin.A_pass == apass {
			auser["aname"] = admin.A_name
			auser["apass"] = admin.A_pass
			token, _ := util.GenerateToken(auser)
			admin.Token = token
			util.Success(c, admin)
			util.DB.Create(&setLogin)
		} else {
			util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
		}
	} else {
		util.Error(c, -1, util.ApiCode.Message[util.ApiCode.FAILED])
	}

}

func (AdminController) TokenVerification(c *gin.Context) {
	mapss := make(map[string]interface{})
	fmt.Println(c.Request.Header.Get("Authorization"))
	token := c.Request.Header.Get("Authorization")
	token = token[6:]
	fmt.Println(util.ConfirmToken(token, mapss))
	auser, _ := util.ConfirmToken(token, mapss)
	util.Success(c, auser)
}
