package sys

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"cmkit/pkg/models"
	"cmkit/pkg/utils"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(endpoints SysEndpoints, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	// 查询公司列表
	listCompanysHandler := kithttp.NewServer(
		endpoints.ListCompanysEndpoint,
		decodeListSearchRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/sys/companys", listCompanysHandler).Methods("GET")

	// 查询部门列表
	listDepartmentsHandler := kithttp.NewServer(
		endpoints.ListDepartmentsEndpoint,
		decodeDeptSearchRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/sys/departments", listDepartmentsHandler).Methods("GET")

	// 添加员工
	addStaffHandler := kithttp.NewServer(
		endpoints.AddStaffEndpoint,
		decodeAddStaffRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/sys/staff", addStaffHandler).Methods("POST")
	// 修改员工
	updateStaffHandler := kithttp.NewServer(
		endpoints.UpdateStaffEndpoint,
		decodeUpdateStaffRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/sys/staff", updateStaffHandler).Methods("PUT")

	// 删除员工
	deleteStaffHandler := kithttp.NewServer(
		endpoints.DeleteStaffEndpoint,
		decodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/sys/staff/{id}", deleteStaffHandler).Methods("DELETE")
	// 查询员工列表
	listStaffsHandler := kithttp.NewServer(
		endpoints.ListStaffsEndpoint,
		decodeListSearchRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/sys/staffs", listStaffsHandler).Methods("GET")

	return r
}

func decodeListSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.URL.Query()
	name := vars.Get("name")
	size := vars.Get("pageSize")
	index := vars.Get("pageIndex")
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		pageSize = 10
	}
	pageIndex, err2 := strconv.Atoi(index)
	if err2 != nil {
		pageIndex = 1
	}

	return map[string]interface{}{"name": name, "pageSize": pageSize, "pageIndex": pageIndex}, nil
}

func decodeDeptSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.URL.Query()
	name := vars.Get("name")
	size := vars.Get("pageSize")
	index := vars.Get("pageIndex")
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		pageSize = 10
	}
	pageIndex, err2 := strconv.Atoi(index)
	if err2 != nil {
		pageIndex = 1
	}

	company := vars.Get("companyId")
	companyID, err3 := strconv.Atoi(company)
	if err3 != nil {
		companyID = 0
	}

	return map[string]interface{}{"name": name, "companyId": companyID, "pageSize": pageSize, "pageIndex": pageIndex}, nil
}

func decodeAddStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var staff models.Staff
	if err := json.NewDecoder(r.Body).Decode(&staff); err != nil {
		return nil, err
	}
	return staff, nil
}

func decodeUpdateStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var staff models.Staff
	if err := json.NewDecoder(r.Body).Decode(&staff); err != nil {
		return nil, err
	}
	return staff, nil
}

func decodeDataIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, utils.ErrBadQueryParams
	}
	staffID, _ := strconv.Atoi(id)
	return models.BaseModel{ID: uint(staffID)}, nil
}
