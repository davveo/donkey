package fake

type LoginAdminUser struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    LoginAdminUserMetaData `json:"data"`
}
type Admin struct {
	AdminID    int    `json:"admin_id"`
	Username   string `json:"username"`
	GroupID    int    `json:"group_id"`
	Nickname   string `json:"nickname"`
	HeadPic    string `json:"head_pic"`
	LastLogin  string `json:"last_login"`
	LastIP     string `json:"last_ip"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
type Token struct {
	GroupID        int    `json:"group_id"`
	Token          string `json:"token"`
	TokenExpires   int    `json:"token_expires"`
	Refresh        string `json:"refresh"`
	RefreshExpires int    `json:"refresh_expires"`
}
type LoginAdminUserMetaData struct {
	Admin Admin `json:"admin"`
	Token Token `json:"token"`
}
