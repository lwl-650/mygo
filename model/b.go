package model

type B struct {
	Bid   int
	Bname string
	Bage  int
	Aaid  int
	A     A
}

func (B) TableName() string {
	return "b"
}
