package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpTypr struct {
}

func (HttpTypr) GetHttp(url string) string {

	response, err := http.Get(url)
	if err != nil {
		//...
	}
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		fmt.Println(string(body))
	}

	return string(body)
}
