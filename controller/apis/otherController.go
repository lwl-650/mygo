package apis

import (
	"mygo/model"
	"mygo/util"

	"github.com/gin-gonic/gin"
)

type OtherController struct{}

type Stu struct {
	Name string `json:"name"`
}

func (OtherController) Ssq(c *gin.Context) {
	// response, err := http.Get("http://apis.juhe.cn/lottery/types?key=b919609752eac679e509cd49f10cfed0")
	// if err != nil {
	// 	//...
	// }
	// defer response.Body.Close() //在回复后必须关闭回复的主体
	// body, err := ioutil.ReadAll(response.Body)
	// if err == nil {
	// 	fmt.Println(string(body))
	// }

	getHttp := util.HttpTypr{}
	res := getHttp.GetHttp("http://apis.juhe.cn/lottery/types?key=b919609752eac679e509cd49f10cfed0")

	util.Success(c, res)
}

func (OtherController) GetSsq(c *gin.Context) {
	getSsqLsit := []model.Ssq{}

	if util.DB.Find(&getSsqLsit).Error != nil {

	} else {
		util.Success(c, getSsqLsit)
	}

}
