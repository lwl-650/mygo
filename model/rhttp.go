package model

type Rhttp struct {
	// gorm.Model
	// `json:"y_id" gorm:"primaryKey"`
	R_id        int     `json:"r_id"`
	R_name      string  `json:"r_name"`
	R_method    string  `json:"r_method"`
	R_url       string  `json:"r_url"`
	R_actualurl string  `json:"r_actualurl"`
	R_status    int     `json:"r_status"`
	R_time      string  `json:"r_time"`
	Adminid     uint    `json:"adminid"`
	Admins      []Admin `gorm:"many2many:rhttp_admin;"`
}

// Admin       Admin  `gorm:"foreignKey:a_id;references:adminid"`
//指定表名
func (Rhttp) TableName() string {
	return "rhttp"
}
