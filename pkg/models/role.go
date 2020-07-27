package models

// Role 用户
type Role struct {
	BaseModel
	Name   string `json:"name" gorm:"size:64"`
	Status int16  `json:"status"` // 0-正常，1-锁定，2-删除
	Remark string `json:"remark"`
}

// TableName role表
func (Role) TableName() string {
	return "t_auth_role"
}
