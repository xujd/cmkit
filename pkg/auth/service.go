package auth

import (
	"cmkit/pkg/models"
	"cmkit/pkg/utils"
	"crypto/sha256"
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
)

// Service 服务接口
type Service interface {
	// 添加用户
	AddUser(user models.User) (string, error)
	// 修改用户
	UpdateUser(user models.User) (string, error)
	// 删除用户
	DeleteUser(id uint) (string, error)
	// 查询用户
	QueryUserByID(id uint) (*models.User, error)
	// 查询用户列表
	ListUsers(name string, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 登录验证
	Login(name, pwd string) (string, error)
	// token续订
	Renewval(token string) (string, error)
	// 获取用户信息
	GetUserInfo(token string) (*models.UserInfo, error)
	// 退出登录
	Logout(token string) (string, error)
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
	if err := s.DB.Where("name = ?", name).First(&user).Error; err != nil {
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
	if err := s.DB.Select("id,created_at,updated_at,deleted_at,name,start_time,end_time,status,remark").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ListUsers 查询用户
func (s AuthService) ListUsers(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&models.User{}) {
		return nil, utils.ErrNotFound
	}
	userdb := s.DB.Model(&models.User{})
	if name != "" {
		userdb = s.DB.Model(&models.User{}).Where("name LIKE ?", "%"+name+"%")
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	userdb.Count(&rowCount)                                            //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var users []models.User
	if err := userdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Select("id,created_at,updated_at,deleted_at,name,start_time,end_time,status,remark").Find(&users).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &users}, nil
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

// GetUserInfo 用户信息
func (s AuthService) GetUserInfo(token string) (*models.UserInfo, error) {
	if token != "" {
		claims, err := ParseToken(token)
		if err != nil {
			return nil, err
		}

		return &models.UserInfo{
			Roles:        []string{"admin"},
			Introduction: "I am a super administrator",
			Avatar:       "./assets/user.gif",
			Name:         claims.Name,
		}, nil
	}
	return nil, utils.ErrBadQueryParams
}

// Logout 退出登录
func (s AuthService) Logout(token string) (string, error) {
	return "success", nil
}
