package home

import (
	"cmkit/pkg/auth"
	"context"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

// HomeEndpoints 首页的Endpoint
type HomeEndpoints struct {
	StatAllResEndpoint      endpoint.Endpoint
	StatSlingByTonEndpoint  endpoint.Endpoint
	GetSlingUsedTopEndpoint endpoint.Endpoint
}

// CreateEndpoints HomeEndpoints
func CreateEndpoints(svc Service) HomeEndpoints {
	statAllResEndpoint := MakeStatAllResEndpoint(svc)
	statAllResEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(statAllResEndpoint)

	statSlingByTonEndpoint := MakeStatSlingByTonEndpoint(svc)
	statSlingByTonEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(statSlingByTonEndpoint)

	getSlingUsedTopEndpoint := MakeGetSlingUsedTopEndpoint(svc)
	getSlingUsedTopEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getSlingUsedTopEndpoint)

	homeEndpoints := HomeEndpoints{
		StatAllResEndpoint:      statAllResEndpoint,
		StatSlingByTonEndpoint:  statSlingByTonEndpoint,
		GetSlingUsedTopEndpoint: getSlingUsedTopEndpoint,
	}

	return homeEndpoints
}

// MakeStatAllResEndpoint 统计各类资源
func MakeStatAllResEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.StatAllRes()

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeStatSlingByTonEndpoint
func MakeStatSlingByTonEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.StatSlingByTon()

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeGetSlingUsedTopEndpoint
func MakeGetSlingUsedTopEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.GetSlingUsedTop()

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
