package admin

type ServiceInterface interface {
	Login(username, password string, lastLogin int64, lastIp, platform string, isGetToken bool) (bool, error)
	Logout() bool
	Close()
}
