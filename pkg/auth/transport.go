package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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
