package util

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func init() {
	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:123456@tcp(127.0.0.1:3306)/vueparent?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
