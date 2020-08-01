package sys

import (
	"cmkit/pkg/auth"
	"cmkit/pkg/models"
	"context"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

// SysEndpoints 系统的Endpoint
type SysEndpoints struct {
	ListCompanysEndpoint    endpoint.Endpoint
	ListDepartmentsEndpoint endpoint.Endpoint
	AddStaffEndpoint        endpoint.Endpoint
	UpdateStaffEndpoint     endpoint.Endpoint
	DeleteStaffEndpoint     endpoint.Endpoint
	ListStaffsEndpoint      endpoint.Endpoint
	ListDictDataEndpoint    endpoint.Endpoint
}

// CreateEndpoints 创建SysEndpoints
func CreateEndpoints(svc Service) SysEndpoints {
	listCompanysEndpoint := MakeListCompanysEndpoint(svc)
	listCompanysEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listCompanysEndpoint)
	listDepartmentsEndpoint := MakeListDepartmentsEndpoint(svc)
	listDepartmentsEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listDepartmentsEndpoint)
	addStaffEndpoint := MakeAddStaffEndpoint(svc)
	addStaffEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(addStaffEndpoint)
	updateStaffEndpoint := MakeUpdateStaffEndpoint(svc)
	updateStaffEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(updateStaffEndpoint)
	deleteStaffEndpoint := MakeDeleteStaffEndpoint(svc)
	deleteStaffEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(deleteStaffEndpoint)
	listStaffsEndpoint := MakeListStaffsEndpoint(svc)
	listStaffsEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listStaffsEndpoint)
	listDictDataEndpoint := MakeListDictDataEndpoint(svc)
	listDictDataEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(listDictDataEndpoint)

	sysEndpoints := SysEndpoints{
		ListCompanysEndpoint:    listCompanysEndpoint,
		ListDepartmentsEndpoint: listDepartmentsEndpoint,
		AddStaffEndpoint:        addStaffEndpoint,
		UpdateStaffEndpoint:     updateStaffEndpoint,
		DeleteStaffEndpoint:     deleteStaffEndpoint,
		ListStaffsEndpoint:      listStaffsEndpoint,
		ListDictDataEndpoint:    listDictDataEndpoint,
	}

	return sysEndpoints
}

// MakeListCompanysEndpoint 查询公司列表
func MakeListCompanysEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListCompanys(req["name"].(string), req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListDepartmentsEndpoint 查询部门列表
func MakeListDepartmentsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListDepartments(req["name"].(string), uint(req["companyId"].(int)), req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeAddStaffEndpoint 添加员工
func MakeAddStaffEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.Staff)

		result, err := svc.AddStaff(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeUpdateStaffEndpoint 修改员工
func MakeUpdateStaffEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.Staff)

		result, err := svc.UpdateStaff(req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeDeleteStaffEndpoint 删除员工
func MakeDeleteStaffEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.BaseModel)

		result, err := svc.DeleteStaff(req.ID)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListStaffsEndpoint 查询员工列表
func MakeListStaffsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListStaffs(req["name"].(string), uint(req["companyId"].(int)), uint(req["departmentId"].(int)), req["pageIndex"].(int), req["pageSize"].(int))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

// MakeListDictDataEndpoint 查询字典
func MakeListDictDataEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(map[string]interface{})
		result, err := svc.ListDict(req["scene"].(string), req["dictType"].(string))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
