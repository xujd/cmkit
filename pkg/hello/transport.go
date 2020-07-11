package hello

import (
	"context"
	"net/http"

	"cmkit/pkg/utils"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(endpoint endpoint.Endpoint, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	helloHandler := kithttp.NewServer(
		endpoint,
		decodeHelloRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/hello/{world}", helloHandler).Methods("GET")

	return r
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	world, ok := vars["world"]
	if !ok {
		return nil, utils.ErrBadQueryParams
	}
	return HelloRequest{World: world}, nil
}
