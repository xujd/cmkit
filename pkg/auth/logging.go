package auth

import (
	"cmkit/pkg/models"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware Make a new type
// that contains Service interface and logger instance
type LoggingMiddleware struct {
	logger log.Logger
	Service
}

// NewLoggingMiddleware 日志
func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &LoggingMiddleware{logger, s}
}

// Login 登录
func (mw LoggingMiddleware) Login(name, pwd string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Login",
			"input", fmt.Sprintf("name=%s, password=%s", name, pwd),
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Login(name, pwd)
	return
}

// Renewval 续订
func (mw LoggingMiddleware) Renewval(oldToken string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Renewval",
			"input", fmt.Sprintf("oldToken=%s", oldToken),
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Renewval(oldToken)
	return
}

// AddUser 添加用户
func (mw LoggingMiddleware) AddUser(user models.User) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddUser",
			"input", fmt.Sprintf("user=%+v", user),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddUser(user)
	return
}

// UpdateUser 修改用户
func (mw LoggingMiddleware) UpdateUser(user models.User) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateUser",
			"input", fmt.Sprintf("user=%+v", user),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateUser(user)
	return
}

// DeleteUser 删除用户
func (mw LoggingMiddleware) DeleteUser(id uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteUser",
			"input", fmt.Sprintf("id=%d", id),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.DeleteUser(id)
	return
}

// QueryUserByID 查询用户
func (mw LoggingMiddleware) QueryUserByID(id uint) (result *models.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "QueryUserByID",
			"input", fmt.Sprintf("id=%d", id),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.QueryUserByID(id)
	return
}

// ListUsers 查询用户列表
func (mw LoggingMiddleware) ListUsers(name string, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListUsers",
			"input", fmt.Sprintf("name=%s, pageIndex=%d, pageSize=%d", name, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListUsers(name, pageIndex, pageSize)
	return
}

// GetUserInfo 获取用户信息
func (mw LoggingMiddleware) GetUserInfo(token string) (result *models.UserInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetUserInfo",
			"input", fmt.Sprintf("token=%s", token),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetUserInfo(token)
	return
}

// Logout 退出登录
func (mw LoggingMiddleware) Logout(token string) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Logout",
			"input", fmt.Sprintf("token=%s", token),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.Logout(token)
	return
}

// AddRole 添加角色
func (mw LoggingMiddleware) AddRole(role models.Role) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddRole",
			"input", fmt.Sprintf("role=%+v", role),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddRole(role)
	return
}

// UpdateRole 修改角色
func (mw LoggingMiddleware) UpdateRole(role models.Role) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateRole",
			"input", fmt.Sprintf("role=%+v", role),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateRole(role)
	return
}

// DeleteRole 删除角色
func (mw LoggingMiddleware) DeleteRole(id uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteRole",
			"input", fmt.Sprintf("id=%d", id),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.DeleteRole(id)
	return
}

// ListRoles 查询角色列表
func (mw LoggingMiddleware) ListRoles(name string, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ListRoles",
			"input", fmt.Sprintf("name=%s, pageIndex=%d, pageSize=%d", name, pageIndex, pageSize),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ListRoles(name, pageIndex, pageSize)
	return
}

// SetUserRole 设置用户角色
func (mw LoggingMiddleware) SetUserRole(userID uint, roleIDs []uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "SetUserRole",
			"input", fmt.Sprintf("userID=%d, roleIDs=%+v", userID, roleIDs),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.SetUserRole(userID, roleIDs)
	return
}

// GetUserRole 获取用户角色
func (mw LoggingMiddleware) GetUserRole(userID uint) (result *[]models.UserRoleRelation, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetUserRole",
			"input", fmt.Sprintf("userID=%d", userID),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetUserRole(userID)
	return
}

// SetRoleFuncs 设置角色权限
func (mw LoggingMiddleware) SetRoleFuncs(roleFunc models.RoleFunc) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "SetRoleFuncs",
			"input", fmt.Sprintf("roleFunc=%+v", roleFunc),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.SetRoleFuncs(roleFunc)
	return
}

// GetRoleFuncs 获取角色权限
func (mw LoggingMiddleware) GetRoleFuncs(roleID uint) (result *models.RoleFunc, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetRoleFuncs",
			"input", fmt.Sprintf("roleID=%d", roleID),
			"result", fmt.Sprintf("%+v", result),
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetRoleFuncs(roleID)
	return
}

// ResetPassword 重置密码
func (mw LoggingMiddleware) ResetPassword(userID uint) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ResetPassword",
			"input", fmt.Sprintf("userID=%d", userID),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ResetPassword(userID)
	return
}

// UpdatePassword 修改密码
func (mw LoggingMiddleware) UpdatePassword(userID uint, password string, newPassword string) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdatePassword",
			"input", fmt.Sprintf("userID=%d, password=%s, newPassword=%s", userID, password, newPassword),
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdatePassword(userID, password, newPassword)
	return
}
