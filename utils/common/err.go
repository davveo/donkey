package common

import "errors"

var (
	ErrNotFound = errors.New("未找到匹配记录")

	ErrAccountNotFound   = errors.New("账号不存在")
	AccountForbbinden    = errors.New("账号已禁用")
	ErrAccountOrPassword = errors.New("账号或密码错误")
)
