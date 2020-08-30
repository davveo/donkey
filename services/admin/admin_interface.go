package admin

import (
	"github.com/davveo/donkey/models/response"
	"time"
)

type ServiceInterface interface {
	Login(username, password string, lastLogin time.Time, lastIp, platform string, isGetToken bool) (*response.AdminResp, error)
	AdminList() (err error, list interface{}, total int)
	Logout() bool
	Close()
}
