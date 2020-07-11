package hello

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// HelloRequest 请求参数
type HelloRequest struct {
	World string
}

// HelloResponse 返回数据
type HelloResponse struct {
	Message string `json:"message"`
}

// MakeHelloEndpoint 创建Hello的Endpoint
func MakeHelloEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HelloRequest)

		message, err := svc.Hello(req.World)

		var resp HelloResponse
		if err != nil {
			return nil, err
		}
		resp = HelloResponse{
			Message: message,
		}

		return resp, nil
	}
}
