package controller

import (
	"net/http"
	"path/filepath"

	"github.com/davveo/donkey/services"
	"github.com/tidwall/gjson"

	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
)

var (
	shortPlatformService = services.AdminService
)

func LoginAdmin(context *gin.Context) {
	//login, err := services.AdminService.Login("admin3")

	result := common.ReadJson(filepath.Join(BaseDir, "data/login.admin.user.json"))
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
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
