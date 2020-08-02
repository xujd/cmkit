package home

import (
	"net/http"

	"github.com/gorilla/mux"

	"cmkit/pkg/utils"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(endpoints HomeEndpoints, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	// 资源统计
	statAllResHandler := kithttp.NewServer(
		endpoints.StatAllResEndpoint,
		utils.DecodeNullRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/home/stat_all_res", statAllResHandler).Methods("GET")

	// 按吨位统计吊索具
	statSlingByTonHandler := kithttp.NewServer(
		endpoints.StatSlingByTonEndpoint,
		utils.DecodeNullRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/home/stat_sling_by_ton", statSlingByTonHandler).Methods("GET")

	// 统计使用次数top
	getSlingUsedTopHandler := kithttp.NewServer(
		endpoints.GetSlingUsedTopEndpoint,
		utils.DecodeNullRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/home/sling_used_top", getSlingUsedTopHandler).Methods("GET")

	// 获取状态统计
	statSlingByStatusHandler := kithttp.NewServer(
		endpoints.StatSlingByStatusEndpoint,
		utils.DecodeNullRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/home/stat_sling_by_status", statSlingByStatusHandler).Methods("GET")
	return r
}
