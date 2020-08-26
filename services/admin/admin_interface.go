package admin

type ServiceInterface interface {
	Login(username string) (bool, error)
	Logout() bool
	Close()
}
