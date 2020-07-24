package utils

import (
	"context"
	"encoding/json"
	"net/http"

	kitjwt "github.com/go-kit/kit/auth/jwt"
)

// EncodeError 对错误信息进行编码
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    codeFormat(err),
		"success": false,
		"message": err.Error(),
	})
}

func codeFormat(err error) int {
	switch err {
	case ErrNotFound:
		return ERR_NOT_FOUND
	case ErrAlreadyExists, ErrUserPwdDismatch, ErrBadQueryParams:
		return ERR_BAD_REQUEST
	case kitjwt.ErrTokenExpired:
		return ERR_TOKEN_EXPIRED
	case kitjwt.ErrTokenInvalid, kitjwt.ErrTokenNotActive:
		return ERR_ILLEGAL_TOKEN
	default:
		return http.StatusInternalServerError
	}
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrUserPwdDismatch, ErrBadQueryParams:
		return http.StatusBadRequest
	case kitjwt.ErrTokenInvalid, kitjwt.ErrTokenExpired, kitjwt.ErrTokenNotActive:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

type errorer interface {
	error() error
}

const (
	// SUCCESS 成功
	SUCCESS int = 20000
	// ERR_NOT_FOUND 未找到
	ERR_NOT_FOUND int = 40000
	// ERR_INTERNAL_SERVER_ERROR 内部错误
	ERR_INTERNAL_SERVER_ERROR int = 40003
	// ERR_BAD_REQUEST 错误请求
	ERR_BAD_REQUEST int = 40005
	// ERR_TOKEN_EXPIRED Token超时
	ERR_TOKEN_EXPIRED int = 50014
	// ERR_ILLEGAL_TOKEN 无效的Token
	ERR_ILLEGAL_TOKEN int = 50008
)

// CmkitResponse 统一返回格式
type CmkitResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// EncodeResponse 对返回信息进行编码
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	cmkitResponse := &CmkitResponse{
		Code:    SUCCESS,
		Success: true,
		Data:    response,
		Message: "",
	}
	return json.NewEncoder(w).Encode(cmkitResponse)
}
