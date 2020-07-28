package models

// User 用户
type User struct {
	BaseModel
	Name      string    `json:"name" gorm:"size:64"`
	Password  string    `json:"password" gorm:"size:64"`
	StartTime *JSONTime `json:"startTime" gorm:"type:timestamp"`
	EndTime   *JSONTime `json:"endTime" gorm:"type:timestamp"`
	Status    int16     `json:"status"` // 0-正常，1-锁定，2-删除
	Remark    string    `json:"remark"`
}

// TableName user表
func (User) TableName() string {
	return "t_auth_user"
}

// UserInfo 用户信息
type UserInfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
	ID           uint     `json:"id"`
}

// UserRoleRelation 用户角色关系
type UserRoleRelation struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"userId"`
	RoleID uint `json:"roleId"`
}

// TableName user_role表
func (UserRoleRelation) TableName() string {
	return "r_auth_user_role"
}
