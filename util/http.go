package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpTypr struct {
}

// 请求第三方url接口
func (HttpTypr) GetHttp(url string) string {

	response, err := http.Get(url)
	if err != nil {
		return ApiCode.Message[ApiCode.FAILED]
	}
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		fmt.Println(string(body))
	}

	return string(body)
}
