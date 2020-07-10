package hello

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(endpoint endpoint.Endpoint, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	helloHandler := kithttp.NewServer(
		endpoint,
		decodeHelloRequest,
		encodeHelloResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/hello/{world}", helloHandler).Methods("GET")

	return r
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	world, ok := vars["world"]
	if !ok {
		return nil, errors.New("bad query params.")
	}
	return HelloRequest{World: world}, nil
}

func encodeHelloResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
