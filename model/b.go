package model

type B struct {
	Id    int
	Bid   int
	Bname string
	As    []A `gorm:"many2many:a_b;"`
}
