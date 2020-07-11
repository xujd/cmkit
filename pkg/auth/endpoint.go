package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// AuthEndpoints
type AuthEndpoints struct {
	LoginEndpoint endpoint.Endpoint
	RenewEndpoint endpoint.Endpoint
}

// AuthRequest
type AuthRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// AuthToken
type AuthToken struct {
	Token string `json:"token"`
}

func MakeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthRequest)

		token, err := svc.Login(req.Name, req.Pwd)

		var resp AuthToken
		if err != nil {
			return nil, err
		} else {
			resp = AuthToken{
				Token: token,
			}
		}

		return resp, nil
	}
}

func MakeRenewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthToken)

		token, err := svc.Renew(req.Token)

		var resp AuthToken
		if err != nil {
			return nil, err
		} else {
			resp = AuthToken{
				Token: token,
			}
		}

		return resp, nil
	}
}
