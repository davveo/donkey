package response

import "github.com/davveo/donkey/models"

//type AdminResp struct {
//	AdminID    uint32 `json:"admin_id"`
//	UserName   string `json:"username"`
//	NickName   string `json:"nickname"`
//	HeadPic    string `json:"head_pic"`
//	LastLogin  string `json:"last_login"`
//	LastIp     string `json:"last_ip"`
//	Status     uint   `json:"status"`
//	CreateTime string `json:"create_time"`
//	UpdateTime string `json:"update_time"`
//	GroupID    uint   `json:"group_id"`
//}

type AdminResp struct {
	Admin *models.Admin
	Token *models.Token
}
