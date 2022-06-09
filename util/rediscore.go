package util

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

var REDIS redis.Conn

func InitRedis() {

	REDIS, err = redis.Dial("tcp", ReadeIni("redis.host"))
	if err != nil {
		log.Fatal("[GIN-MYSQL] connect to redis error:" + err.Error())
	}
	log.Println("[GIN-Redis] connected success")
	// defer REDIS.Close()
}

func CloseRedis() {
	fmt.Println("关闭是否执行")
	REDIS.Close()
}
func SETredis(key string, val interface{}) {
	_, err = REDIS.Do("SET", key, val)
	if err != nil {
		fmt.Println("redis set failed:", err)
	} else {
		fmt.Println("redis set ok")
	}

}

// func CloseRedis() {
// 	REDIS.Close()
// }
// func redis_connect() redis.Conn {

// 	c, err := redis.Dial("tcp", ReadeIni("redis.host"))
// 	if err != nil {
// 		log.Fatal("[GIN-MYSQL] connect to redis error:" + err.Error())
// 	}
// 	log.Println("[GIN-Redis] connected success")
// 	return c
// }
