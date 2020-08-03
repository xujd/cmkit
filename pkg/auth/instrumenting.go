package auth

import (
	"cmkit/pkg/models"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware Instrumenting
type InstrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingMiddleware New
func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &InstrumentingMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

// Login 登录
func (s *InstrumentingMiddleware) Login(name, pwd string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Login").Add(1)
		s.requestLatency.With("method", "Login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Login(name, pwd)
}

// Renewval 续订
func (s *InstrumentingMiddleware) Renewval(oldToken string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Renewval").Add(1)
		s.requestLatency.With("method", "Renewval").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Renewval(oldToken)
}

// AddUser 添加用户
func (s *InstrumentingMiddleware) AddUser(user models.User) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AddUser").Add(1)
		s.requestLatency.With("method", "AddUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddUser(user)
}

// UpdateUser 修改用户
func (s *InstrumentingMiddleware) UpdateUser(user models.User) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateUser").Add(1)
		s.requestLatency.With("method", "UpdateUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdateUser(user)
}

// DeleteUser 删除用户
func (s *InstrumentingMiddleware) DeleteUser(id uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteUser").Add(1)
		s.requestLatency.With("method", "DeleteUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteUser(id)
}

// QueryUserByID 查询用户
func (s InstrumentingMiddleware) QueryUserByID(id uint) (result *models.User, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "QueryUserByID").Add(1)
		s.requestLatency.With("method", "QueryUserByID").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.QueryUserByID(id)
}

// ListUsers 查询用户列表
func (s InstrumentingMiddleware) ListUsers(name string, pageIndex int, pageSize int) (result *models.SearchResult, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListUsers").Add(1)
		s.requestLatency.With("method", "ListUsers").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListUsers(name, pageIndex, pageSize)
}

// GetUserInfo 查询用户信息
func (s InstrumentingMiddleware) GetUserInfo(token string) (result *models.UserInfo, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetUserInfo").Add(1)
		s.requestLatency.With("method", "GetUserInfo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetUserInfo(token)
}

// Logout 退出登录
func (s InstrumentingMiddleware) Logout(token string) (result string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Logout").Add(1)
		s.requestLatency.With("method", "Logout").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Logout(token)
}

// AddRole 添加角色
func (s InstrumentingMiddleware) AddRole(role models.Role) (result string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AddRole").Add(1)
		s.requestLatency.With("method", "AddRole").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AddRole(role)
}

// UpdateRole 修改角色
func (s InstrumentingMiddleware) UpdateRole(role models.Role) (result string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateRole").Add(1)
		s.requestLatency.With("method", "UpdateRole").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdateRole(role)
}

// DeleteRole 删除角色
func (s InstrumentingMiddleware) DeleteRole(id uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "DeleteRole").Add(1)
		s.requestLatency.With("method", "DeleteRole").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteRole(id)
}

// ListRoles 查询角色列表
func (s InstrumentingMiddleware) ListRoles(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListRoles").Add(1)
		s.requestLatency.With("method", "ListRoles").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ListRoles(name, pageIndex, pageSize)
}

// SetUserRole 设置用户角色
func (s InstrumentingMiddleware) SetUserRole(userID uint, roleIDs []uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SetUserRole").Add(1)
		s.requestLatency.With("method", "SetUserRole").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.SetUserRole(userID, roleIDs)
}

// GetUserRole 获取用户角色
func (s InstrumentingMiddleware) GetUserRole(userID uint) (*[]models.UserRoleRelation, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetUserRole").Add(1)
		s.requestLatency.With("method", "GetUserRole").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetUserRole(userID)
}

// SetRoleFuncs 设置角色权限
func (s InstrumentingMiddleware) SetRoleFuncs(roleFunc models.RoleFunc) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SetRoleFuncs").Add(1)
		s.requestLatency.With("method", "SetRoleFuncs").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.SetRoleFuncs(roleFunc)
}

// GetRoleFuncs 获取角色权限
func (s InstrumentingMiddleware) GetRoleFuncs(roleID uint) (*models.RoleFunc, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetRoleFuncs").Add(1)
		s.requestLatency.With("method", "GetRoleFuncs").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetRoleFuncs(roleID)
}

// ResetPassword 重置密码
func (s InstrumentingMiddleware) ResetPassword(userID uint) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ResetPassword").Add(1)
		s.requestLatency.With("method", "ResetPassword").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ResetPassword(userID)
}

// UpdatePassword 修改密码
func (s InstrumentingMiddleware) UpdatePassword(userID uint, password string, newPassword string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdatePassword").Add(1)
		s.requestLatency.With("method", "UpdatePassword").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.UpdatePassword(userID, password, newPassword)
}
