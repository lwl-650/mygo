package model

type A struct {
	Aid   int
	Aname string
	Aage  int
}

func (A) TableName() string {
	return "a"
}
