package controller

import (
	"fmt"
	"time"

	"github.com/davveo/donkey/utils/log"

	"github.com/allegro/bigcache"

	"github.com/davveo/donkey/utils/captcha"
	"github.com/davveo/donkey/utils/response"
	"github.com/gin-gonic/gin"
)

var (
	config = bigcache.Config{
		Shards:             16,              // 存储的条目数量，值必须是2的幂
		LifeWindow:         2 * time.Second, // 超时后条目被删除
		MaxEntriesInWindow: 0,               // 在 Life Window 中的最大数量
		MaxEntrySize:       0,               // 条目最大尺寸，以字节为单位
		HardMaxCacheSize:   0,               // 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}
	Cache, _ = bigcache.NewBigCache(config)
)

func Captcha(context *gin.Context) {
	// captchaType=[audio, string, math, chinese]
	captchaType := context.DefaultQuery("captchaType", "")

	id, b64s, err := captcha.GenerateCaptcha(captchaType)
	if err != nil {
		log.ERROR.Println(err)
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithData(
		map[string]interface{}{
			"data":      b64s,
			"captchaId": id,
		}, context)

}

func HealthCheck(context *gin.Context) {
	response.Ok(context)
}

func QrCode(context *gin.Context) {
	value, exists := context.Get("client")
	fmt.Println(value, exists)
	//clientType := context.DefaultQuery("client_type", "-1")
	//appKey := ""
	//resp := map[string]interface{}{"captcha": true, "session_id": ""}
	//app := models.App{}
	//appResult, _ := app.FindByKey(nil, appKey)
	//if appResult != nil {
	//	if appResult.Captcha == 0 {
	//		context.JSON(http.StatusBadRequest, gin.H{
	//			"status":  http.StatusBadRequest,
	//			"message": "验证码已经禁用!",
	//			"data":    resp,
	//		})
	//		return
	//	}
	//}
	//
	//if clientType == "-1" {
	//	resp["session_id"] = ""
	//} else {
	//	resp["session_id"] = ""
	//}
	//
	//context.JSON(http.StatusBadRequest, gin.H{
	//	"status":  http.StatusBadRequest,
	//	"message": err.Error(),
	//	"data":    resp,
	//})
}

func GetClientToken(context *gin.Context) string {
	value, exists := context.Get("client")
	fmt.Println(value, exists)
	return ""
}
