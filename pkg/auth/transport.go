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
		decodeUserIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/user/{id}", deleteUserHandler).Methods("DELETE")

	// 查询用户
	queryUserByIDHandler := kithttp.NewServer(
		endpoints.QueryUserByIDEndpoint,
		decodeUserIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/user/{id}", queryUserByIDHandler).Methods("GET")
	// 查询用户
	listUsersHandler := kithttp.NewServer(
		endpoints.ListUsersEndpoint,
		decodeRenewvalRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/auth/users", listUsersHandler).Methods("GET")
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

func decodeUserIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, utils.ErrBadQueryParams
	}
	userID, _ := strconv.Atoi(id)
	return models.BaseModel{ID: uint(userID)}, nil
}
