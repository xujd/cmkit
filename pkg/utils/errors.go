package utils

import "errors"

var (
	// ErrAlreadyExists 记录已存在
	ErrAlreadyExists = errors.New("记录已存在")
	// ErrNotFound 未找到
	ErrNotFound = errors.New("未找到")
	// ErrUserPwdDismatch 用户名或密码错误
	ErrUserPwdDismatch = errors.New("用户名或密码错误")
	// ErrBadQueryParams 查询参数错误
	ErrBadQueryParams = errors.New("查询参数错误")
)
