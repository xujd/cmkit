package res

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"cmkit/pkg/utils"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the handling service.
func MakeHandler(endpoints ResEndpoints, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(utils.EncodeError),
	}

	// 查询吊索具列表
	listSlingsHandler := kithttp.NewServer(
		endpoints.ListSlingsEndpoint,
		decodeSlingListSearchRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/slings", listSlingsHandler).Methods("GET")
	// 添加吊索具
	addSlingHandler := kithttp.NewServer(
		endpoints.AddSlingEndpoint,
		decodeSlingRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/sling", addSlingHandler).Methods("POST")
	// 修改吊索具
	updateSlingHandler := kithttp.NewServer(
		endpoints.UpdateSlingEndpoint,
		decodeSlingRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/sling", updateSlingHandler).Methods("PUT")

	// 删除吊索具
	deleteSlingHandler := kithttp.NewServer(
		endpoints.DeleteSlingEndpoint,
		utils.DecodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/sling/{id}", deleteSlingHandler).Methods("DELETE")

	// 查询智能柜列表
	listCabinetsHandler := kithttp.NewServer(
		endpoints.ListCabinetsEndpoint,
		decodeCabinetListSearchRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/cabinets", listCabinetsHandler).Methods("GET")
	// 添加智能柜
	addCabinetHandler := kithttp.NewServer(
		endpoints.AddCabinetEndpoint,
		decodeCabinetRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/cabinet", addCabinetHandler).Methods("POST")
	// 修改智能柜
	updateCabinetHandler := kithttp.NewServer(
		endpoints.UpdateCabinetEndpoint,
		decodeCabinetRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/cabinet", updateCabinetHandler).Methods("PUT")

	// 删除智能柜
	deleteCabinetHandler := kithttp.NewServer(
		endpoints.DeleteCabinetEndpoint,
		utils.DecodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/cabinet/{id}", deleteCabinetHandler).Methods("DELETE")

	// 查询箱格
	listGridsHandler := kithttp.NewServer(
		endpoints.ListCabinetGridsEndpoint,
		utils.DecodeDataIDRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/cabinet_grids/{id}", listGridsHandler).Methods("GET")

	// 存
	storeHandler := kithttp.NewServer(
		endpoints.StoreEndpoint,
		utils.DecodeCommonRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/store", storeHandler).Methods("POST")
	// 存取
	takeReturnHandler := kithttp.NewServer(
		endpoints.TakeReturnEndpoint,
		utils.DecodeCommonRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/take_return", takeReturnHandler).Methods("POST")
	// 存取
	takeReturnByResIDHandler := kithttp.NewServer(
		endpoints.TakeReturnByResIDEndpoint,
		decodeTakeReturnRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/take_return_by_res", takeReturnByResIDHandler).Methods("POST")

	// 查询使用日志
	getResUseLogHandler := kithttp.NewServer(
		endpoints.GetTakeReturnLogEndpoint,
		decodeResUseLogSearchRequest,
		utils.EncodeResponse,
		append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))...,
	)

	r.Handle("/res/uselog", getResUseLogHandler).Methods("GET")
	return r
}

func decodeSlingListSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.URL.Query()
	name := vars.Get("name")
	slingType := vars.Get("slingType")
	maxTonnage := vars.Get("maxTonnage")
	useStatus := vars.Get("useStatus")
	inspectStatus := vars.Get("inspectStatus")
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

	slingTypeID, err3 := strconv.Atoi(slingType)
	if err3 != nil {
		slingTypeID = 0
	}

	maxTonnageID, err4 := strconv.Atoi(maxTonnage)
	if err4 != nil {
		maxTonnageID = 0
	}

	useStatusID, err4 := strconv.Atoi(useStatus)
	if err4 != nil {
		useStatusID = 0
	}

	inspectStatusID, err5 := strconv.Atoi(inspectStatus)
	if err5 != nil {
		inspectStatusID = 0
	}

	return map[string]interface{}{
		"name":          name,
		"slingType":     slingTypeID,
		"maxTonnage":    maxTonnageID,
		"useStatus":     useStatusID,
		"inspectStatus": inspectStatusID,
		"pageSize":      pageSize,
		"pageIndex":     pageIndex,
	}, nil
}

func decodeSlingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var sling Sling
	if err := json.NewDecoder(r.Body).Decode(&sling); err != nil {
		return nil, err
	}
	return sling, nil
}

func decodeCabinetListSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

	return map[string]interface{}{
		"name":      name,
		"pageSize":  pageSize,
		"pageIndex": pageIndex,
	}, nil
}

func decodeCabinetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var cabinet Cabinet
	if err := json.NewDecoder(r.Body).Decode(&cabinet); err != nil {
		return nil, err
	}
	return cabinet, nil
}

func decodeTakeReturnRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var useLog UseLog
	if err := json.NewDecoder(r.Body).Decode(&useLog); err != nil {
		return nil, err
	}
	return useLog, nil
}

func decodeResUseLogSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.URL.Query()
	resName := vars.Get("resName")
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
	flag := vars.Get("returnFlag")
	returnFlag, err3 := strconv.Atoi(flag)
	if err3 != nil {
		returnFlag = 0
	}
	takeStaff := vars.Get("takeStaff")
	takeStaffID, err4 := strconv.Atoi(takeStaff)
	if err4 != nil {
		takeStaffID = 0
	}
	returnStaff := vars.Get("returnStaff")
	returnStaffID, err4 := strconv.Atoi(returnStaff)
	if err4 != nil {
		returnStaffID = 0
	}

	return map[string]interface{}{
		"resName":       resName,
		"pageSize":      pageSize,
		"takeStaff":     takeStaffID,
		"returnStaff":   returnStaffID,
		"takeStartTime": vars.Get("takeStartTime"),
		"takeEndTime":   vars.Get("takeEndTime"),
		"returnFlag":    returnFlag,
		"pageIndex":     pageIndex,
	}, nil
}
