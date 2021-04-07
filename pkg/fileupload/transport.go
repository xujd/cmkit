package fileupload

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"cmkit/pkg/utils"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(endpoints FileUploadEndpoints, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	// 文件上传
	listSlingsHandler := kithttp.NewServer(
		endpoints.FileUploadEndpoint,
		decodeUploadRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/file/upload", listSlingsHandler).Methods("POST")

	return r
}

func decodeUploadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
