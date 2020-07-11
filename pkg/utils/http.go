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
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	})
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

// CmkitResponse 统一返回格式
type CmkitResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
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
		Success: true,
		Data:    response,
		Error:   "",
	}
	return json.NewEncoder(w).Encode(cmkitResponse)
}
