package model

type Admin struct {
	// gorm.Model
	A_id       int     `json:"a_id"`
	A_name     string  `json:"a_name"`
	A_pass     string  `json:"a_pass"`
	A_portrait string  `json:"a_portrait"`
	A_grade    int     `json:"a_grade"`
	Rhttps     []Rhttp `gorm:"many2many:rhttp_admin;"`
	Aname      string  `json:"aname"`
	Aid   int
}

// 创建关联表名函数
func (Admin) TableName() string {
	// 返回表名
	return "admin"
}
