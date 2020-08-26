package controller

import (
	"net/http"
	"path/filepath"

	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func Navigation(context *gin.Context) {
	result := common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json"))
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
}
