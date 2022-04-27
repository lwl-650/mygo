package model

type R_http struct {
	R_id        int    `json:"r_id"`
	R_name      string `json:"r_name"`
	R_method    string `json:"r_method"`
	R_url       string `json:"r_url"`
	R_actualurl string `json:"r_actualurl"`
	R_status    int    `json:"r_status"`
	R_time      string `json:"r_time"`
}

//指定表名
func (R_http) TableName() string {
	return "r_http"
}
