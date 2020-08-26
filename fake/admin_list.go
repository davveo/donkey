package fake

type AdminList struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			AdminID      int    `json:"admin_id"`
			Username     string `json:"username"`
			GroupID      int    `json:"group_id"`
			Nickname     string `json:"nickname"`
			HeadPic      string `json:"head_pic"`
			LastLogin    string `json:"last_login"`
			LastIP       string `json:"last_ip"`
			Status       int    `json:"status"`
			CreateTime   string `json:"create_time"`
			UpdateTime   string `json:"update_time"`
			GetAuthGroup struct {
				Name   string `json:"name"`
				Status int    `json:"status"`
			} `json:"get_auth_group"`
			LastIPRegion string `json:"last_ip_region"`
		} `json:"items"`
		TotalResult int `json:"total_result"`
	} `json:"data"`
}
