package admin

import "github.com/davveo/donkey/models/response"

type ServiceInterface interface {
	Login(username, password string, lastLogin int64, lastIp, platform string, isGetToken bool) (*response.AdminResp, error)
	Logout() bool
	Close()
}
