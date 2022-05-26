package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.AddConfigPath("./conf/") // 文件所在目录
	config.SetConfigName("conf")    // 文件名
	config.SetConfigType("ini")     // 文件类型

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	host := config.GetString("mysql.host") // 读取配置
	fmt.Println("viper load ini: ", host)
	// b := model.Book{}
	// fmt.Println(b)
	// b.SetBook("zs", 18)
	// fmt.Println(b)

	// b.Siy()
	// b.Siy2()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// 创建一个默认的路由引擎

	// util.Tack()

	// r := gin.Default()
	// r.Use(router.Content) //全局中间件
	// router.IndexRoutersInit(r)
	// router.HostRouters(r)
	// // 启动HTTP服务，默认在0.0.0.0:8080启动服务
	// r.Run(":8088")

	// book := util.Book{}
	// book.SetAge(18)
	// book.SetName("zs")
	// fmt.Println(book.GetAge())
	// fmt.Println(book.GetName())
	// fmt.Println(book)
}
