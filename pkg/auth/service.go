package auth

import (
	"cmkit/pkg/models"
	"cmkit/pkg/utils"
	"crypto/sha256"
	"fmt"
	"math"
	"strconv"
	"time"

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
	// 添加角色
	AddRole(role models.Role) (string, error)
	// 修改角色
	UpdateRole(role models.Role) (string, error)
	// 删除角色
	DeleteRole(id uint) (string, error)
	// 查询角色列表
	ListRoles(name string, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 设置用户角色
	SetUserRole(userID uint, roleIDs []uint) (string, error)
	// 获取用户角色
	GetUserRole(userID uint) (*[]models.UserRoleRelation, error)
	// 设置角色权限
	SetRoleFuncs(roleFunc models.RoleFunc) (string, error)
	// 获取角色权限
	GetRoleFuncs(roleID uint) (*models.RoleFunc, error)
	// 重置密码
	ResetPassword(userID uint) (string, error)
	// 修改密码
	UpdatePassword(userID uint, password string, newPassword string) (string, error)
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
	// 默认密码
	user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte("123456a?"+user.Name)))
	// 用户名不能为空
	if user.Name == "" {
		return "", utils.ErrUserNameIsNull
	}
	// 员工未指定
	if user.StaffID == 0 {
		return "", utils.ErrUserStaffIsNull
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
	// 默认用户不准修改
	if user.ID == 1 {
		return "", utils.ErrNoUpdate
	}
	if !s.DB.HasTable(&models.User{}) {
		if err := s.DB.CreateTable(&models.User{}).Error; err != nil {
			return "", err
		}
	}
	_, err0 := s.QueryUserByID(user.ID)
	if err0 != nil {
		return "", utils.ErrUserNotFound
	}
	data := map[string]interface{}{
		"Remark": user.Remark,
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
	// 根用户不准删除
	if id == 1 {
		return "", utils.ErrNoDelete
	}
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
	selectStr := "t_auth_user.id,t_auth_user.created_at,t_auth_user.updated_at,t_auth_user.deleted_at,t_auth_user.name,t_auth_user.start_time,t_auth_user.end_time,t_auth_user.status,t_auth_user.remark,t_auth_user.staff_id, t_sys_staff.name AS staff_name"

	if err := s.DB.Table("t_auth_user").Select(selectStr).
		Joins("JOIN t_sys_staff ON t_auth_user.staff_id = t_sys_staff.id").
		Where("t_auth_user.deleted_at IS NULL AND t_auth_user.id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ListUsers 查询用户
func (s AuthService) ListUsers(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&models.User{}) {
		return nil, utils.ErrNotFound
	}
	selectStr := "t_auth_user.id,t_auth_user.created_at,t_auth_user.updated_at,t_auth_user.deleted_at,t_auth_user.name,t_auth_user.start_time,t_auth_user.end_time,t_auth_user.status,t_auth_user.remark,t_auth_user.staff_id, t_sys_staff.name AS staff_name"
	userdb := s.DB.Table("t_auth_user").Select(selectStr).
		Joins("JOIN t_sys_staff ON t_auth_user.staff_id = t_sys_staff.id").
		Where("t_auth_user.deleted_at IS NULL")

	if name != "" {
		userdb = userdb.Where("t_auth_user.name LIKE ?", "%"+name+"%")
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
	if err := userdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
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

	timeFormatStr := "2006-01-02 15:04:05"
	// 开始生效时间
	if user.StartTime != nil {
		startTime := (*time.Time)(user.StartTime).Format(timeFormatStr)
		t1, _ := time.ParseInLocation(timeFormatStr, string(startTime), time.Local)
		if t1.After(time.Now()) {
			return "", utils.ErrUserNotEffective
		}
	}
	// 结束生效时间
	if user.EndTime != nil {
		endTime := (*time.Time)(user.EndTime).Format(timeFormatStr)
		t2, _ := time.ParseInLocation(timeFormatStr, string(endTime), time.Local)
		if t2.Before(time.Now()) {
			return "", utils.ErrUserExpired
		}
	}
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(pwd+name)))
	if user.Name == name && user.Password == password {
		token, err := Sign(name, strconv.Itoa(int(user.ID)))
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
		id, _ := strconv.Atoi(claims.UserId)
		user, err := s.QueryUserByID(uint(id))
		if err != nil {
			return nil, utils.ErrUserNotFound
		}

		userInfo := &models.UserInfo{
			Introduction: user.Remark,
			Avatar:       "./assets/user.gif",
			Name:         user.Name,
			ID:           user.ID,
			StaffName:    user.StaffName,
		}

		var roles []models.Role
		s.DB.Raw("SELECT * FROM t_auth_role WHERE id IN (SELECT role_id FROM r_auth_user_role WHERE user_id=?)", user.ID).Scan(&roles)

		userInfo.Roles = make([]string, len(roles))
		for key, value := range roles {
			userInfo.Roles[key] = value.Name
		}

		return userInfo, nil
	}
	return nil, utils.ErrBadQueryParams
}

// Logout 退出登录
func (s AuthService) Logout(token string) (string, error) {
	return "success", nil
}

// AddRole 添加角色
func (s AuthService) AddRole(role models.Role) (string, error) {
	if !s.DB.HasTable(&models.Role{}) {
		if err := s.DB.CreateTable(&models.Role{}).Error; err != nil {
			return "", err
		}
	}
	// 角色名称不能为空
	if role.Name == "" {
		return "", utils.ErrRoleNameIsNull
	}

	role0, _ := s.QueryRoleByName(role.Name)
	if role0 != nil {
		return "", utils.ErrRoleAlreadyExists
	}

	if err := s.DB.Create(&role).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// QueryRoleByName 查询角色
func (s AuthService) QueryRoleByName(name string) (*models.Role, error) {
	if !s.DB.HasTable(&models.Role{}) {
		return nil, utils.ErrNotFound
	}
	var role models.Role
	if err := s.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

// QueryRoleByID 查询角色
func (s AuthService) QueryRoleByID(id uint) (*models.Role, error) {
	if !s.DB.HasTable(&models.Role{}) {
		return nil, utils.ErrNotFound
	}
	var role models.Role
	if err := s.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

// UpdateRole 修改角色
func (s AuthService) UpdateRole(role models.Role) (string, error) {
	// 默认角色不准修改
	if role.ID == 1 {
		return "", utils.ErrNoUpdate
	}
	if !s.DB.HasTable(&models.Role{}) {
		if err := s.DB.CreateTable(&models.Role{}).Error; err != nil {
			return "", err
		}
	}
	_, err0 := s.QueryRoleByID(role.ID)
	if err0 != nil {
		return "", utils.ErrRoleNotFound
	}
	if err := s.DB.Save(&role).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// DeleteRole 删除角色
func (s AuthService) DeleteRole(id uint) (string, error) {
	// 根角色不准删除
	if id == 1 {
		return "", utils.ErrNoDelete
	}
	if err := s.DB.Where("id = ?", id).Delete(&models.Role{}).Error; err != nil {
		return "", nil
	}
	return "success", nil
}

// ListRoles 获取角色列表
func (s AuthService) ListRoles(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&models.Role{}) {
		return nil, utils.ErrNotFound
	}
	roledb := s.DB.Model(&models.Role{})
	if name != "" {
		roledb = s.DB.Model(&models.Role{}).Where("name LIKE ?", "%"+name+"%")
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	roledb.Count(&rowCount)                                            //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var roles []models.Role
	if err := roledb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&roles).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &roles}, nil
}

// SetUserRole 设置用户角色
func (s AuthService) SetUserRole(userID uint, roleIDs []uint) (string, error) {
	// 检查表是否存在
	if !s.DB.HasTable(&models.UserRoleRelation{}) {
		if err := s.DB.CreateTable(&models.UserRoleRelation{}).Error; err != nil {
			return "", err
		}
	}
	// 事务
	tx := s.DB.Begin()
	// 先删除旧数据
	if err := tx.Where("user_id = ?", userID).Delete(&models.UserRoleRelation{}).Error; err != nil {
		return "", nil
	}
	// 增加新关系
	for _, value := range roleIDs {
		if err := tx.Create(&models.UserRoleRelation{UserID: userID, RoleID: value}).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}
	tx.Commit()
	return "success", nil
}

// GetUserRole 查询用户角色
func (s AuthService) GetUserRole(userID uint) (*[]models.UserRoleRelation, error) {
	// 检查表是否存在
	if !s.DB.HasTable(&models.UserRoleRelation{}) {
		return nil, utils.ErrNotFound
	}
	fmt.Println("hello")
	var userRoles []models.UserRoleRelation
	if err := s.DB.Model(&models.UserRoleRelation{}).Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	return &userRoles, nil
}

// SetRoleFuncs 设置角色权限
func (s AuthService) SetRoleFuncs(roleFunc models.RoleFunc) (string, error) {
	// 检查表是否存在
	if !s.DB.HasTable(&models.RoleFunc{}) {
		if err := s.DB.CreateTable(&models.RoleFunc{}).Error; err != nil {
			return "", err
		}
	}
	// 事务
	tx := s.DB.Begin()
	// 先删除旧数据
	if err := tx.Where("role_id = ?", roleFunc.ID).Delete(&models.RoleFunc{}).Error; err != nil {
		return "", nil
	}
	// 增加新关系
	if err := tx.Create(&roleFunc).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return "success", nil
}

// GetRoleFuncs 获取角色权限
func (s AuthService) GetRoleFuncs(roleID uint) (*models.RoleFunc, error) {
	// 检查表是否存在
	if !s.DB.HasTable(&models.RoleFunc{}) {
		return nil, utils.ErrNotFound
	}
	var roleFunc models.RoleFunc
	if err := s.DB.Where("role_id = ?", roleID).First(&roleFunc).Error; err != nil {
		return nil, err
	}

	return &roleFunc, nil
}

// ResetPassword 重置密码
func (s AuthService) ResetPassword(userID uint) (string, error) {
	// 默认用户不准修改
	if userID == 1 {
		return "", utils.ErrNoUpdate
	}
	if !s.DB.HasTable(&models.User{}) {
		if err := s.DB.CreateTable(&models.User{}).Error; err != nil {
			return "", err
		}
	}
	user0, err0 := s.QueryUserByID(userID)
	if err0 != nil {
		return "", utils.ErrUserNotFound
	}
	// 两次加密
	password1 := fmt.Sprintf("%x", sha256.Sum256([]byte("123456a?"+user0.Name)))
	data := map[string]interface{}{
		"Password": fmt.Sprintf("%x", sha256.Sum256([]byte(password1+user0.Name))),
	}
	if err := s.DB.Model(&user0).Updates(data).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// UpdatePassword 修改密码
func (s AuthService) UpdatePassword(userID uint, password string, newPassword string) (string, error) {
	// 默认用户不准修改
	if userID == 1 {
		return "", utils.ErrNoUpdate
	}
	if !s.DB.HasTable(&models.User{}) {
		return "", utils.ErrUserNotFound
	}
	// 查询用户是否存在
	var user models.User
	if err := s.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return "", err
	}
	// 不存在
	if user.ID == 0 {
		return "", utils.ErrUserNotFound
	}
	// 确认密码
	oldPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password+user.Name)))
	if oldPassword != user.Password {
		return "", utils.ErrPwdDismatch
	}
	data := map[string]interface{}{
		"Password": fmt.Sprintf("%x", sha256.Sum256([]byte(newPassword+user.Name))),
	}
	if err := s.DB.Model(&user).Updates(data).Error; err != nil {
		return "", err
	}
	return "success", nil
}
