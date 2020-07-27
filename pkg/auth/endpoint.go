package auth

import (
	"cmkit/pkg/models"
	"context"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

// AuthEndpoints 权限的Endpoint
type AuthEndpoints struct {
	LoginEndpoint         endpoint.Endpoint
	RenewvalEndpoint      endpoint.Endpoint
	AddUserEndpoint       endpoint.Endpoint
	UpdateUserEndpoint    endpoint.Endpoint
	DeleteUserEndpoint    endpoint.Endpoint
	QueryUserByIDEndpoint endpoint.Endpoint
	ListUsersEndpoint     endpoint.Endpoint
	GetUserInfoEndpoint   endpoint.Endpoint
	LogoutEndpoint        endpoint.Endpoint
	AddRoleEndpoint       endpoint.Endpoint
	UpdateRoleEndpoint    endpoint.Endpoint
	DeleteRoleEndpoint    endpoint.Endpoint
	ListRolesEndpoint     endpoint.Endpoint
	SetUserRoleEndpoint   endpoint.Endpoint
}

// AuthRequest 登录请求
type AuthRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// AuthToken Token
type AuthToken struct {
	Token string `json:"token"`
}

// CreateEndpoints 创建AuthEndpoints
func CreateEndpoints(svc Service) AuthEndpoints {
	loginEndpoint := MakeLoginEndpoint(svc)
	renewvalEndpoint := MakeRenewvalEndpoint(svc)
	renewvalEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(renewvalEndpoint)
	logoutEndpoint := MakeLogoutEndpoint(svc)
	logoutEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(logoutEndpoint)

	addUserEndpoint := MakeAddUserEndpoint(svc)
	addUserEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(addUserEndpoint)
	updateUserEndpoint := MakeUpdateUserEndpoint(svc)
	updateUserEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(updateUserEndpoint)
	deleteUserEndpoint := MakeDeleteUserEndpoint(svc)
	deleteUserEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(deleteUserEndpoint)
	queryUserByIDEndpoint := MakeQueryUserByIDEndpoint(svc)
	queryUserByIDEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(queryUserByIDEndpoint)
	listUsersEndpoint := MakeListUsersEndpoint(svc)
	listUsersEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listUsersEndpoint)
	getUserInfoEndpoint := MakeGetUserInfoEndpoint(svc)
	getUserInfoEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getUserInfoEndpoint)

	addRoleEndpoint := MakeAddRoleEndpoint(svc)
	addRoleEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(addRoleEndpoint)
	updateRoleEndpoint := MakeUpdateRoleEndpoint(svc)
	updateRoleEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(updateRoleEndpoint)
	deleteRoleEndpoint := MakeDeleteRoleEndpoint(svc)
	deleteRoleEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(deleteRoleEndpoint)
	listRolesEndpoint := MakeListRolesEndpoint(svc)
	listRolesEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listRolesEndpoint)

	setUserRoleEndpoint := MakeSetUserRoleEndpoint(svc)
	setUserRoleEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(setUserRoleEndpoint)

	authEndpoints := AuthEndpoints{
		LoginEndpoint:         loginEndpoint,
		RenewvalEndpoint:      renewvalEndpoint,
		AddUserEndpoint:       addUserEndpoint,
		UpdateUserEndpoint:    updateUserEndpoint,
		DeleteUserEndpoint:    deleteUserEndpoint,
		QueryUserByIDEndpoint: queryUserByIDEndpoint,
		ListUsersEndpoint:     listUsersEndpoint,
		GetUserInfoEndpoint:   getUserInfoEndpoint,
		LogoutEndpoint:        logoutEndpoint,
		AddRoleEndpoint:       addRoleEndpoint,
		UpdateRoleEndpoint:    updateRoleEndpoint,
		DeleteRoleEndpoint:    deleteRoleEndpoint,
		ListRolesEndpoint:     listRolesEndpoint,
		SetUserRoleEndpoint:   setUserRoleEndpoint,
	}

	return authEndpoints
}

// MakeLoginEndpoint 创建登录Endpoint
func MakeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthRequest)

		token, err := svc.Login(req.UserName, req.Password)

		var resp AuthToken
		if err != nil {
			return nil, err
		}
		resp = AuthToken{
			Token: token,
		}

		return resp, nil
	}
}

// MakeRenewvalEndpoint 创建续订Endpoint
func MakeRenewvalEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthToken)

		token, err := svc.Renewval(req.Token)

		var resp AuthToken
		if err != nil {
			return nil, err
		}
		resp = AuthToken{
			Token: token,
		}

		return resp, nil
	}
}

// MakeAddUserEndpoint 创建用户
func MakeAddUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.User)

		result, err := svc.AddUser(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeUpdateUserEndpoint 修改用户
func MakeUpdateUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.User)

		result, err := svc.UpdateUser(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeDeleteUserEndpoint 删除用户
func MakeDeleteUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.DeleteUser(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeQueryUserByIDEndpoint 查询用户
func MakeQueryUserByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.QueryUserByID(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListUsersEndpoint 查询用户列表
func MakeListUsersEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListUsers(req["name"].(string), req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeGetUserInfoEndpoint 获取用户信息
func MakeGetUserInfoEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthToken)
		data, err := svc.GetUserInfo(req.Token)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

// MakeLogoutEndpoint 退出登录
func MakeLogoutEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthToken)
		data, err := svc.Logout(req.Token)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

// MakeAddRoleEndpoint 添加角色
func MakeAddRoleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.Role)

		result, err := svc.AddRole(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeUpdateRoleEndpoint 修改角色
func MakeUpdateRoleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.Role)

		result, err := svc.UpdateRole(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeDeleteRoleEndpoint 删除角色
func MakeDeleteRoleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.DeleteRole(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListRolesEndpoint 查询用户列表
func MakeListRolesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListRoles(req["name"].(string), req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeSetUserRoleEndpoint 设置用户角色
func MakeSetUserRoleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		roleIds := make([]uint, len(req["roleIds"].([]interface{})))
		for i, value := range req["roleIds"].([]interface{}) {
			roleIds[i] = uint(value.(float64))
		}
		result, err := svc.SetUserRole(uint(req["userId"].(float64)), roleIds)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
