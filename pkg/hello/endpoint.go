package hello

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// HelloRequest
type HelloRequest struct {
	World string
}

// HelloResponse
type HelloResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func MakeHelloEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HelloRequest)

		message, err := svc.Hello(req.World)

		var resp HelloResponse
		if err != nil {
			resp = HelloResponse{
				Success: err == nil,
				Message: message,
				Error:   err.Error(),
			}
		} else {
			resp = HelloResponse{
				Success: err == nil,
				Message: message,
			}
		}

		return resp, nil
	}
}
