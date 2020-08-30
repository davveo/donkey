package enum

const (
	// 系统默认权限
	AUTH_SUPER_ADMINISTRATOR = 1
	AUTH_ADMINISTRATOR       = 2
	AUTH_CLIENT              = 3
	AUTH_GUEST               = 4
)

const (
	// -1:游客 0:顾客 1:管理组
	VISITOR = -1
	CUSTOM  = 0
	ADMIN   = 1
)
