package model

type Admin struct {
	A_id       uint    `json:"a_id"`
	A_name     string  `json:"a_name"`
	A_pass     string  `json:"a_pass"`
	A_portrait string  `json:"a_portrait"`
	A_grade    int     `json:"a_grade"`
	Rhttp      []Rhttp `gorm:"many2many:rhttp_admin;"`
}

// 创建关联表名函数
func (Admin) TableName() string {
	// 返回表名
	return "admin"
}
