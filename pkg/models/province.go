package models

// Province 省份
type Province struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name" gorm:"size:32"` // 省名称
	Status int16  `json:"status"`              // 状态：0-正常，1-停用
	Remark string `json:"remark"`              // 说明
}

// TableName province
func (Province) TableName() string {
	return "t_sys_province"
}
