package util

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}

func Snow() int64 {
	// 使用Format格式化日期
	fmt.Println(time.Now().Format("2006-01-02"))
	if err := Init(time.Now().Format("2006-01-02"), 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return 0
	}

	id := GenID()
	fmt.Println(id)
	return id
}
