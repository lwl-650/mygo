package util

import (
	"time"
)

// 三元运算符
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

// func If(condition bool, trueVal, falseVal interface{}) interface{} {
// 	if condition {
// 		return trueVal
// 	} else {
// 		return falseVal
// 	}
// }

// 格式化时间
func GetTime() string {
	getTime := time.Now()
	return getTime.Format("2006-01-02 15:04:05")
}
