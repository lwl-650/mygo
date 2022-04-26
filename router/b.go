package router

import (
	"fmt"
	"mygo/util"

	"github.com/robfig/cron/v3"
)

func fn() {
	getHttp := util.HttpTypr{}
	res := getHttp.GetHttp("http://apis.juhe.cn/lottery/types?key=b919609752eac679e509cd49f10cfed0")
	// util.Success(c, res)
	fmt.Println(res)
}

func main() {
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	task := func() {
		fn()
	}
	//定时任务
	spec := "*/5 * * * * ?" //cron表达式，每五秒一次
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} //阻塞主线程停止
}
