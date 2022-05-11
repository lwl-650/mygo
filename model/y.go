package model

type Y struct {
	// `json:"y_id" gorm:"primaryKey"`
	Yid   int `json:"y_id" gorm:"primaryKey"`
	Yname string
	X     []X `gorm:"many2many:x_y" json:"gety"`
}
