package apis

import (
	"fmt"
	"mygo/util"

	"github.com/gin-gonic/gin"
)

type TesController struct {
}

func (TesController) GetArr(c *gin.Context) {
	arr := []int{1, 2, 3}
	util.Success(c, arr)
}

func (TesController) GetMap(c *gin.Context) {
	// arr := []int{1, 2, 3}
	dict := make(map[string]int)
	dict["France"] = 0
	dict["Italy"] = 1
	dict["Japan"] = 2
	dict["India "] = 3

	util.Success(c, dict)
}
func (TesController) GetMapArr(c *gin.Context) {
	dict := make([]map[string]int, 2)
	dict[0] = make(map[string]int)
	dict[1] = make(map[string]int)
	dict[0]["France"] = 0
	dict[0]["Italy"] = 1
	dict[1]["Japan"] = 2
	dict[1]["India "] = 3
	util.Success(c, dict)
}

func (TesController) GetArrMap(c *gin.Context) {
	dict := make(map[string]interface{})

	arrint := []int{1, 2, 3}
	var numbers []bool
	numbers = append(numbers, true)
	numbers = append(numbers, false)
	arrstring := []interface{}{"1", "2", numbers}
	arr := []interface{}{arrint, arrstring, 3, numbers}

	dict2 := make([]map[string]int, 2)
	dict2[0] = make(map[string]int)
	dict2[1] = make(map[string]int)
	dict2[0]["France"] = 0
	dict2[0]["Italy"] = 1
	dict2[1]["Japan"] = 2
	dict2[1]["India "] = 3

	map1 := make(map[string]interface{})
	map2 := make(map[string]interface{})
	map3 := make(map[string]interface{})
	map1["map2"] = map2
	map2["map3"] = map3

	dict["data"] = arr
	dict["name"] = 4
	dict["dict2"] = dict2
	dict["map1"] = map1

	fmt.Println(dict)
	util.Success(c, dict)
}
