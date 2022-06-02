package model

type Xue struct {
	Id  int64  `json:"id"`
	Xue string `json:"xue"`
}

//指定表名
func (Xue) TableName() string {
	return "xue"
}
