package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/davveo/donkey/utils/log"

	"github.com/allegro/bigcache"

	"github.com/davveo/donkey/fake"
	"github.com/davveo/donkey/utils/captcha"
	"github.com/davveo/donkey/utils/common"
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

}

func AppInstallUpdated(context *gin.Context) {
	var appInstall fake.AppInstall
	result := common.ReadJson(filepath.Join(BaseDir, "data/app.install.json"))
	err := json.Unmarshal([]byte(result), &appInstall)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  appInstall.Status,
		"message": appInstall.Message,
		"data":    appInstall.Data,
	})
}
