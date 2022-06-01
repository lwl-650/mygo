package model

type Ssq struct {
	Ssqid   int    `json:"ssqid" gorm:"primaryKey"`
	Ssqtext string `json:"ssqtext"`
}

func (Ssq) TableName() string {
	return "ssq"
}
