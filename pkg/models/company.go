package models

// Company 公司
type Company struct {
	BaseModel
	Name   string `json:"name" gorm:"size:128"` // 公司名称
	Status int16  `json:"status"`               // 状态：0-正常，1-停用
	Remark string `json:"remark"`               // 说明
}

// TableName company表
func (Company) TableName() string {
	return "auth_company"
}
