package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// AuthEndpoints 权限的Endpoint
type AuthEndpoints struct {
	LoginEndpoint endpoint.Endpoint
	RenewEndpoint endpoint.Endpoint
}

// AuthRequest 登录请求
type AuthRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// AuthToken Token
type AuthToken struct {
	Token string `json:"token"`
}

// MakeLoginEndpoint 创建登录Endpoint
func MakeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthRequest)

		token, err := svc.Login(req.Name, req.Pwd)

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

// MakeRenewEndpoint 创建续订Endpoint
func MakeRenewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthToken)

		token, err := svc.Renew(req.Token)

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
