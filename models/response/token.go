package response

type TokenResp struct {
	GroupID        uint   `json:"group_id"`
	Token          string `json:"token"`
	TokenExpires   string `json:"token_expires"`
	Refresh        string `json:"refresh"`
	RefreshExpires string `json:"refresh_expires"`
}
