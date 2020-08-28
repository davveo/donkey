package controller

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/davveo/donkey/services"
	"github.com/tidwall/gjson"

	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
)

var (
	shortPlatformService = services.AdminService
)

func LoginAdmin(context *gin.Context) {
	var (
		status  int
		message string
		data    interface{}
	)
	// TODO 获取外部请求参数
	resp, err := services.AdminService.Login(
		"admin2", "admin2",
		time.Now().Unix(), "127.0.0.1",
		"admin", true)
	if err != nil {
		message = err.Error()
		data = nil
		status = http.StatusBadRequest
	} else {
		message = "success"
		data = map[string]interface{}{
			"admin": resp.Admin,
			"token": resp.Token,
		}
		status = http.StatusOK
	}
	context.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func AdminItem(context *gin.Context) {

}

func CheckAdmin(context *gin.Context) {

}

func AdminList(context *gin.Context) {
	result := common.ReadJson(filepath.Join(BaseDir, "data/admin.list.json"))
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
}
