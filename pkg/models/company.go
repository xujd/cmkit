package models

// Company 公司
type Company struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name" gorm:"size:128"` // 公司名称
	Status int16  `json:"status"`               // 状态：0-正常，1-停用
	Remark string `json:"remark"`               // 说明
}

// TableName company表
func (Company) TableName() string {
	return "t_sys_company"
}
