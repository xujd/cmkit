package models

// User 用户
type User struct {
	BaseModel
	Name      string    `json:"name" gorm:"size:64"`
	Password  string    `json:"password" gorm:"size:64"`
	StartTime *JSONTime `json:"startTime"`
	EndTime   *JSONTime `json:"endTime"`
	Status    int16     `json:"status"` // 0-正常，1-锁定，2-删除
	Remark    string    `json:"remark"`
	Salt      string    `json:"-"`
}

// TableName user表
func (User) TableName() string {
	return "auth_user"
}

// UserInfo 用户信息
type UserInfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}
