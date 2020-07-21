package utils

import "errors"

var (
	// ErrAlreadyExists 记录已存在
	ErrAlreadyExists = errors.New("记录已存在")
	// ErrUserAlreadyExists 用户已存在
	ErrUserAlreadyExists = errors.New("用户已存在")
	// ErrUserNotFound 用户不存在
	ErrUserNotFound = errors.New("用户不存在")
	// ErrNameOrPasswordIsNull 用户名或密码为空
	ErrNameOrPasswordIsNull = errors.New("用户名或密码为空")
	// ErrNotFound 未找到
	ErrNotFound = errors.New("未找到")
	// ErrUserPwdDismatch 用户名或密码错误
	ErrUserPwdDismatch = errors.New("用户名或密码错误")
	// ErrUserStatus 用户状态异常
	ErrUserStatus = errors.New("用户状态异常")
	// ErrBadQueryParams 查询参数错误
	ErrBadQueryParams = errors.New("查询参数错误")
)
