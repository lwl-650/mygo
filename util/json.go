package util

import (
	"encoding/json"
)

// 字符串转JSON
func JSON(str string) map[string]interface{} {
	// 字符串
	// str := `{"name":"test","product_id":"1","number":"110011","price":"0.01","is_on_sale":"true"}`
	var p interface{}
	json.Unmarshal([]byte(str), &p)
	// 现在我们需要从这个interface{}解析出里面的数据
	m := p.(map[string]interface{})
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Printf("%s is string, value: %s\n", k, vv)
	// 	case int:
	// 		fmt.Printf("%s is int, value: %d\n", k, vv)
	// 	case int64:
	// 		fmt.Printf("%s is int64, value: %d\n", k, vv)
	// 	case bool:
	// 		fmt.Printf("%s is bool, vaule: %v", k, vv)
	// 	default:
	// 		fmt.Printf("%s is unknow type\n", k)
	// 	}
	// }
	return m
}
