package model

type Loginadmin struct {
	Lid   int    `json:"lid"`
	Lname string `json:"lname"`
	Ltime string `json:"ltime"`
}

// 创建关联表名函数
func (Loginadmin) TableName() string {
	// 返回表名
	return "loginadmin"
}
