package model

type X struct {
	Xid   int `gorm:"primaryKey" json:"xxid"`
	Xname string
	Y     []Y `gorm:"many2many:x_y" json:"getx"`
}
