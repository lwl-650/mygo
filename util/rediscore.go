package util

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type RedisCore struct {
	local     string
	RedisPool redis.Conn
}

func (red *RedisCore) init() {
	//通过go 向 redis 写入数据 和 读取数据
	//1. 连接到 redis

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial err = ", err)
		return
	}
	red.RedisPool = conn
	defer conn.Close() //关闭连接
	//fmt.Println("conn success ...", conn)
}

func (r *RedisCore) Setredis(k, v string) error {

	_, err = r.RedisPool.Do("SET", k, v)
	defer r.RedisPool.Close()
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}

	// c := r.RedisPool.Get()
	// defer c.Close()
	// _, err := c.Do("SET", k, v)
	return err
}
