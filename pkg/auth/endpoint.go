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
	logoutEndpoint := MakeLogoutEndpoint(svc)
	logoutEndpoint = kitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(logoutEndpoint)

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
