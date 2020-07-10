package auth

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

// AuthRequest
type AuthRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// AuthResponse
type AuthResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

func MakeAuthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthRequest)

		token, err := svc.Login(req.Name, req.Pwd)

		var resp AuthResponse
		if err != nil {
			resp = AuthResponse{
				Success: err == nil,
				Token:   fmt.Sprintf("Bearer %s", token),
				Error:   err.Error(),
			}
		} else {
			resp = AuthResponse{
				Success: err == nil,
				Token:   fmt.Sprintf("Bearer %s", token),
			}
		}

		return resp, nil
	}
}
