package auth

import (
	"cmkit/pkg/models"
	"cmkit/pkg/utils"
	"crypto/sha256"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Service 服务接口
type Service interface {
	AddUser(user models.User) (string, error)
	UpdateUser(user models.User) (string, error)
	DeleteUser(id uint) (string, error)
	QueryUserByID(id uint) (*models.User, error)
	ListUsers() (*[]models.User, error)
	Login(name, pwd string) (string, error)
	Renewval(token string) (string, error)
}

// AuthService 权限服务
type AuthService struct {
	DB *gorm.DB
}

// AddUser 添加用户
func (s AuthService) AddUser(user models.User) (string, error) {
	if !s.DB.HasTable(&models.User{}) {
		if err := s.DB.CreateTable(&models.User{}).Error; err != nil {
			return "", err
		}
	}
	// 用户名和密码不能为空
	if user.Name == "" || user.Password == "" {
		return "", utils.ErrNameOrPasswordIsNull
	}

	user0, _ := s.QueryUserByName(user.Name)
	if user0 != nil {
		return "", utils.ErrUserAlreadyExists
	}

	user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password+user.Name)))
	if err := s.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// UpdateUser 修改用户
func (s AuthService) UpdateUser(user models.User) (string, error) {
	if !s.DB.HasTable(&models.User{}) {
		if err := s.DB.CreateTable(&models.User{}).Error; err != nil {
			return "", err
		}
	}
	user0, err0 := s.QueryUserByID(user.ID)
	if err0 != nil {
		return "", utils.ErrUserNotFound
	}
	data := map[string]interface{}{
		"Remark": user.Remark,
	}
	if user.Password != "" {
		data["Password"] = fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password+user0.Name)))
	}
	if user.StartTime != nil {
		data["StartTime"] = user.StartTime
	}
	if user.EndTime != nil {
		data["EndTime"] = user.EndTime
	}
	if user.Status > -1 {
		data["Status"] = user.Status
	}
	if err := s.DB.Model(&user).Updates(data).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// DeleteUser 删除用户
func (s AuthService) DeleteUser(id uint) (string, error) {
	if err := s.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return "", nil
	}
	return "success", nil
}

// QueryUserByName 查询用户
func (s AuthService) QueryUserByName(name string) (*models.User, error) {
	if !s.DB.HasTable(&models.User{}) {
		return nil, utils.ErrNotFound
	}
	var user models.User
	if err := s.DB.Unscoped().Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// QueryUserByID 查询用户
func (s AuthService) QueryUserByID(id uint) (*models.User, error) {
	if !s.DB.HasTable(&models.User{}) {
		return nil, utils.ErrNotFound
	}
	var user models.User
	if err := s.DB.Unscoped().Select("id,created_at,updated_at,deleted_at,name,start_time,end_time,status,remark").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ListUsers 查询用户
func (s AuthService) ListUsers() (*[]models.User, error) {
	if !s.DB.HasTable(&models.User{}) {
		return nil, utils.ErrNotFound
	}
	var users []models.User
	if err := s.DB.Unscoped().Select("id,created_at,updated_at,deleted_at,name,start_time,end_time,status,remark").Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

// Login 登录
func (s AuthService) Login(name, pwd string) (string, error) {
	user, err := s.QueryUserByName(name)
	if err != nil {
		return "", utils.ErrUserPwdDismatch
	}
	// 用户状态异常
	if user.Status != 0 {
		return "", utils.ErrUserStatus
	}
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(pwd+name)))
	if user.Name == name && user.Password == password {
		token, err := Sign(name, pwd)
		return token, err
	}

	return "", utils.ErrUserPwdDismatch
}

// Renewval Token续订
func (s AuthService) Renewval(oldToken string) (string, error) {
	if oldToken != "" {
		token, err := Resign(oldToken)
		return token, err
	}

	return "", utils.ErrBadQueryParams
}
