package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name      string `gorm:"size:64"`
	PassWd    string `gorm:"size:64"`
	StartTime time.Time
	EndTime   time.Time
	Status    int16
	Remark    string
}

// TableName user表
func (User) TableName() string {
	return "auth_user"
}
