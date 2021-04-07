package fileupload

import (
	"cmkit/pkg/auth"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

// FileUploadEndpoints 的Endpoint
type FileUploadEndpoints struct {
	FileUploadEndpoint endpoint.Endpoint
}

// CreateEndpoints FileUploadEndpoints
func CreateEndpoints(svc Service) FileUploadEndpoints {
	uploadEndpoint := MakeUploadEndpoint(svc)
	uploadEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(uploadEndpoint)
	uploadEndpoints := FileUploadEndpoints{
		FileUploadEndpoint: uploadEndpoint,
	}

	return uploadEndpoints
}

func MakeUploadEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.Upload(request.(*http.Request))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
