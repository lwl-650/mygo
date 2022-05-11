package model

type A struct {
	Id    int
	Aid   int
	Aname string
	Bs    []B `gorm:"many2many:a_b;"`
}
