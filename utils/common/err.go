package common

import "errors"

var (
	ErrNotFound          = errors.New("未找到匹配记录")
	ErrAccountNotFound   = errors.New("账号不存在")
	AccountForbbinden    = errors.New("账号已禁用")
	ErrAccountOrPassword = errors.New("账号或密码错误")
)

const (
	NO  = 0 // 禁止, 停用
	YES = 1 // 允许, 开通
)
