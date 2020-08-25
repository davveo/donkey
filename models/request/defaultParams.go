package request

type DefaultParams struct {
	Method    string `json:"method"  binding:"required"`
	Token     string `json:"token" binding:"required"`
	AppKey    string `json:"appkey" binding:"required"`
	TimeStamp int64  `json:"timestamp" binding:"required"`
	Format    string `json:"format" binding:"required"`
	Sign      string `json:"sign" binding:"required"`
}
