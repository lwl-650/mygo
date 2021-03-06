package apis

import (
	"fmt"
	"mygo/model"
	"mygo/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type users struct {
	Name string      `json:"name"`
	Age  int         `json:"age"`
	List interface{} `json:"list"`
}

func (UserController) Index(c *gin.Context) {
	// c.JSON：返回JSON格式的数据
	// users.Name = "zs"
	// users.Age = 16
	// fmt.Println(users)
	// c.JSON(200, gin.H{
	// 	"message": users,
	// })
	// c.JSON(http.StatusOK, users)
	var obj struct {
		A    string      `json:"A"`
		B    int         `json:"B"`
		Arrs interface{} `json:"arrs"`
	}
	var k = []string{"sd", "Sd", "ds"}
	obj.A = "AAAA"
	obj.B = 9999
	obj.Arrs = k
	getuser := users{}
	getuser.Name = "lisi"
	getuser.Age = 2222
	getuser.List = obj

	util.Success(c, getuser)
}

func (UserController) Add(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"name": "zsw",
	// 	"age":  16,
	// })
	fmt.Println(c.Request.Header.Get("token"))
	util.Error(c, -1, "出了一个打错")
}

func (UserController) FindUser(c *gin.Context) {
	userList := []model.User{}
	page, _ := strconv.Atoi(c.PostForm("page"))
	num, _ := strconv.Atoi(c.PostForm("num"))
	getPage := util.If(page > 1, (page-1)*5, 0)
	num = util.If(num == 0, 10, num)
	util.DB.Limit(num).Offset(getPage).Find(&userList)
	util.Success(c, userList)
}

//通过id查询
func (UserController) FindById(c *gin.Context) {
	user := model.User{}
	sid := c.PostForm("id")
	if sid != "" {
		util.DB.First(&user, sid)
		util.Other(c, int(util.ApiCode.AUTHFORMATERROR),
			util.ApiCode.GetMessage(util.ApiCode.CREATEUSERFAILED), user)
	} else {
		util.Error(c, -1, "操作失败")
	}
}

func (UserController) AddUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	avatar := c.PostForm("avatar")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	fmt.Println(name, password, avatar, gender, city)
	if name != "" && password != "" {
		user := model.User{Name: name, Password: password,
			Avatar: avatar, Gender: gender, City: city,
		}
		if err := util.DB.Create(&user).Error; err != nil {
			fmt.Println("插入失败", err)
			util.Error(c, -1, "添加失败")
			return
		} else {
			util.Success(c, err)
		}
	} else {
		util.Error(c, -1, "不能为空。")
	}
	// result := db.Create(&user) // 通过数据的指针来创建
}

func (UserController) DelById(c *gin.Context) {
	id := c.PostForm("id")
	fmt.Println(id)
	if id != "" && util.DB.Delete(&model.User{}, id).Error == nil {
		util.Success(c, 1)
	} else {
		util.Error(c, -1, "删除失败")
	}

}

func (UserController) Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	user := model.User{}
	if util.DB.Where("name = ?", name).First(&user).Error == nil && name == user.Name && password == user.Password {
		util.Success(c, user)
	} else {
		util.Error(c, -1, "loginErr")
	}

}
