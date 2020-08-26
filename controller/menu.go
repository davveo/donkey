package controller

import (
	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"net/http"
	"os"
	"path/filepath"
)

var (
	BaseDir, _ = os.Getwd()
)

func MenuAuthList(context *gin.Context) {
	result := common.ReadJson(filepath.Join(BaseDir, "data/get.menu.auth.list.json"))
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
}

func MenuList(context *gin.Context) {
	result := common.ReadJson(filepath.Join(BaseDir, "data/menu.list.json"))
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
}
