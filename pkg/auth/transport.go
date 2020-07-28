package auth

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
func MakeHandler(endpoints AuthEndpoints, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	loginHandler := kithttp.NewServer(
		endpoints.LoginEndpoint,
		decodeLoginRequest,
		utils.EncodeResponse,
		opts...,
	)

	r.Handle("/auth/login", loginHandler).Methods("POST")

	renewvalHandler := kithttp.NewServer(
		endpoints.RenewvalEndpoint,
		decodeRenewvalRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/renewval", renewvalHandler).Methods("GET")

	// 添加用户
	addUserHandler := kithttp.NewServer(
		endpoints.AddUserEndpoint,
		decodeAddUserRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/user", addUserHandler).Methods("POST")
	// 修改用户
	updateUserHandler := kithttp.NewServer(
		endpoints.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/user", updateUserHandler).Methods("PUT")

	// 删除用户
	deleteUserHandler := kithttp.NewServer(
		endpoints.DeleteUserEndpoint,
		decodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/user/{id}", deleteUserHandler).Methods("DELETE")

	// 查询用户
	queryUserByIDHandler := kithttp.NewServer(
		endpoints.QueryUserByIDEndpoint,
		decodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/user/{id}", queryUserByIDHandler).Methods("GET")
	// 查询用户列表
	listUsersHandler := kithttp.NewServer(
		endpoints.ListUsersEndpoint,
		decodeUserListRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/users", listUsersHandler).Methods("GET")
	// 查询用户信息
	getUserInfoHandler := kithttp.NewServer(
		endpoints.GetUserInfoEndpoint,
		decodeRenewvalRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/userinfo", getUserInfoHandler).Methods("GET")

	// 退出登录
	logoutHandler := kithttp.NewServer(
		endpoints.LogoutEndpoint,
		decodeRenewvalRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/logout", logoutHandler).Methods("GET")

	// 添加角色
	addRoleHandler := kithttp.NewServer(
		endpoints.AddRoleEndpoint,
		decodeAddRoleRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/role", addRoleHandler).Methods("POST")
	// 修改角色
	updateRoleHandler := kithttp.NewServer(
		endpoints.UpdateRoleEndpoint,
		decodeUpdateRoleRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/role", updateRoleHandler).Methods("PUT")

	// 删除角色
	deleteRoleHandler := kithttp.NewServer(
		endpoints.DeleteRoleEndpoint,
		decodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/role/{id}", deleteRoleHandler).Methods("DELETE")

	// 查询角色列表
	listRolesHandler := kithttp.NewServer(
		endpoints.ListRolesEndpoint,
		decodeUserListRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/roles", listRolesHandler).Methods("GET")

	// 设置用户角色
	setUserRoleHandler := kithttp.NewServer(
		endpoints.SetUserRoleEndpoint,
		decodeUserRoleRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/userrole", setUserRoleHandler).Methods("POST")

	// 设置角色权限
	setRoleFuncsHandler := kithttp.NewServer(
		endpoints.SetRoleFuncsEndpoint,
		decodeRoleFuncRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/rolefunc", setRoleFuncsHandler).Methods("POST")

	// 查询角色权限
	getRoleFuncsHandler := kithttp.NewServer(
		endpoints.GetRoleFuncsEndpoint,
		decodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/rolefunc/{id}", getRoleFuncsHandler).Methods("GET")

	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var loginRequest AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		return nil, err
	}
	return loginRequest, nil
}

func decodeRenewvalRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := AuthToken{
		Token: r.Header["Authorization"][0],
	}
	return req, nil
}

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func decodeDataIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, utils.ErrBadQueryParams
	}
	userID, _ := strconv.Atoi(id)
	return models.BaseModel{ID: uint(userID)}, nil
}

func decodeUserListRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func decodeAddRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return nil, err
	}
	return role, nil
}

func decodeUpdateRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return nil, err
	}
	return role, nil
}

func decodeUserRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func decodeRoleFuncRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var roleFunc models.RoleFunc
	if err := json.NewDecoder(r.Body).Decode(&roleFunc); err != nil {
		return nil, err
	}
	return roleFunc, nil
}
