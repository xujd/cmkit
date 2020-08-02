package res

import (
	"cmkit/pkg/auth"
	"cmkit/pkg/models"
	"context"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

// ResEndpoints 资源的Endpoint
type ResEndpoints struct {
	ListSlingsEndpoint  endpoint.Endpoint
	AddSlingEndpoint    endpoint.Endpoint
	UpdateSlingEndpoint endpoint.Endpoint
	DeleteSlingEndpoint endpoint.Endpoint

	ListCabinetsEndpoint      endpoint.Endpoint
	AddCabinetEndpoint        endpoint.Endpoint
	UpdateCabinetEndpoint     endpoint.Endpoint
	DeleteCabinetEndpoint     endpoint.Endpoint
	ListCabinetGridsEndpoint  endpoint.Endpoint
	StoreEndpoint             endpoint.Endpoint
	TakeReturnEndpoint        endpoint.Endpoint
	TakeReturnByResIDEndpoint endpoint.Endpoint
	GetTakeReturnLogEndpoint  endpoint.Endpoint
}

// CreateEndpoints ResEndpoints
func CreateEndpoints(svc Service) ResEndpoints {
	listSlingsEndpoint := MakeListSlingsEndpoint(svc)
	listSlingsEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listSlingsEndpoint)
	addSlingEndpoint := MakeAddSlingEndpoint(svc)
	addSlingEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(addSlingEndpoint)
	updateSlingEndpoint := MakeUpdateSlingEndpoint(svc)
	updateSlingEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(updateSlingEndpoint)
	deleteSlingEndpoint := MakeDeleteSlingEndpoint(svc)
	deleteSlingEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(deleteSlingEndpoint)

	listCabinetsEndpoint := MakeListCabinetsEndpoint(svc)
	listCabinetsEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listCabinetsEndpoint)
	addCabinetEndpoint := MakeAddCabinetEndpoint(svc)
	addCabinetEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(addCabinetEndpoint)
	updateCabinetEndpoint := MakeUpdateCabinetEndpoint(svc)
	updateCabinetEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(updateCabinetEndpoint)
	deleteCabinetEndpoint := MakeDeleteCabinetEndpoint(svc)
	deleteCabinetEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(deleteCabinetEndpoint)

	listCabinetGridsEndpoint := MakeListCabinetGridsEndpoint(svc)
	listCabinetGridsEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listCabinetGridsEndpoint)
	storeEndpoint := MakeStoreEndpoint(svc)
	storeEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(storeEndpoint)
	takeReturnEndpoint := MakeTakeReturnEndpoint(svc)
	takeReturnEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(takeReturnEndpoint)
	takeReturnByResIDEndpoint := MakeTakeReturnByResIDEndpoint(svc)
	takeReturnByResIDEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(takeReturnByResIDEndpoint)

	getTakeReturnLogEndpoint := MakeGetTakeReturnLogEndpoint(svc)
	getTakeReturnLogEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getTakeReturnLogEndpoint)

	resEndpoints := ResEndpoints{
		ListSlingsEndpoint:  listSlingsEndpoint,
		AddSlingEndpoint:    addSlingEndpoint,
		UpdateSlingEndpoint: updateSlingEndpoint,
		DeleteSlingEndpoint: deleteSlingEndpoint,

		ListCabinetsEndpoint:      listCabinetsEndpoint,
		AddCabinetEndpoint:        addCabinetEndpoint,
		UpdateCabinetEndpoint:     updateCabinetEndpoint,
		DeleteCabinetEndpoint:     deleteCabinetEndpoint,
		ListCabinetGridsEndpoint:  listCabinetGridsEndpoint,
		StoreEndpoint:             storeEndpoint,
		TakeReturnEndpoint:        takeReturnEndpoint,
		TakeReturnByResIDEndpoint: takeReturnByResIDEndpoint,
		GetTakeReturnLogEndpoint:  getTakeReturnLogEndpoint,
	}

	return resEndpoints
}

// MakeListSlingsEndpoint 查询吊索具列表
func MakeListSlingsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListSlings(req["name"].(string),
			uint(req["slingType"].(int)), uint(req["maxTonnage"].(int)),
			uint(req["useStatus"].(int)), uint(req["inspectStatus"].(int)),
			req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeAddSlingEndpoint 添加吊索具
func MakeAddSlingEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Sling)

		result, err := svc.AddSling(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeUpdateSlingEndpoint 修改吊索具
func MakeUpdateSlingEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Sling)

		result, err := svc.UpdateSling(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeDeleteSlingEndpoint 删除吊索具
func MakeDeleteSlingEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.DeleteSling(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListCabinetsEndpoint 查询智能柜列表
func MakeListCabinetsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListCabinets(req["name"].(string),
			req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeAddCabinetEndpoint 添加智能柜
func MakeAddCabinetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Cabinet)

		result, err := svc.AddCabinet(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeUpdateCabinetEndpoint 修改智能柜
func MakeUpdateCabinetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Cabinet)

		result, err := svc.UpdateCabinet(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeDeleteCabinetEndpoint 删除智能柜
func MakeDeleteCabinetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.DeleteCabinet(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListCabinetGridsEndpoint 查询箱格
func MakeListCabinetGridsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.ListCabinetGrids(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeStoreEndpoint 存
func MakeStoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.Store(uint(req["cabinetId"].(int)),
			uint(req["gridNo"].(int)), uint(req["resId"].(int)))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeTakeReturnEndpoint 取还
func MakeTakeReturnEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.TakeReturn(uint(req["cabinetId"].(float64)),
			uint(req["gridNo"].(float64)), int(req["flag"].(float64)))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeTakeReturnByResIDEndpoint 取还
func MakeTakeReturnByResIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UseLog)
		result, err := svc.TakeReturnByResID(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeGetTakeReturnLogEndpoint 查询取还日志
func MakeGetTakeReturnLogEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})

		result, err := svc.GetTakeReturnLog(req["resName"].(string),
			uint(req["takeStaff"].(int)), uint(req["returnStaff"].(int)),
			req["takeStartTime"].(string), req["takeEndTime"].(string),
			req["returnFlag"].(int),
			req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
