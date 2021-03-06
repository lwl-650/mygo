package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	Openid   string `json:"openid"`
	City     string `json:"city"`
}

//指定表名
func (User) TableName() string {
	return "user"
}
