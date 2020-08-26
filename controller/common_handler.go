package controller

import (
	"net/http"
	"path/filepath"

	"github.com/davveo/donkey/models/request"
	"github.com/davveo/donkey/utils/common"
	"github.com/davveo/donkey/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var MaptoMethod = map[string]string{
	"defalut":                   "",
	"get.setting.list":          common.ReadJson(filepath.Join(BaseDir, "data/setting.list.json")),
	"get.upload.module":         common.ReadJson(filepath.Join(BaseDir, "data/upload.module.json")),
	"get.action.log.list":       common.ReadJson(filepath.Join(BaseDir, "data/action.log.list.json")),
	"get.help.router":           common.ReadJson(filepath.Join(BaseDir, "data/help.router.json")),
	"get.auth.rule.list":        common.ReadJson(filepath.Join(BaseDir, "data/auth.rule.list.json")),
	"get.auth.group.list":       common.ReadJson(filepath.Join(BaseDir, "data/auth.group.list.json")),
	"query.app.install.updated": common.ReadJson(filepath.Join(BaseDir, "data/app.install.json")),
	"get.menu.auth.list":        common.ReadJson(filepath.Join(BaseDir, "data/get.menu.auth.list.json")),
	"get.menu.list":             common.ReadJson(filepath.Join(BaseDir, "data/menu.list.json")),
	"add.admin.item":            common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json")),
	"check.admin.nickname":      common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json")),
	"login.admin.user":          common.ReadJson(filepath.Join(BaseDir, "data/login.admin.user.json")),
	"get.admin.list":            common.ReadJson(filepath.Join(BaseDir, "data/admin.list.json")),
	"get.message.user.unread":   common.ReadJson(filepath.Join(BaseDir, "data/message.user.unread.json")),
	"get.navigation.list":       common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json")),
	"get.notice.tpl.list":       common.ReadJson(filepath.Join(BaseDir, "data/notice.tpl.list.json")),
	"get.payment.list":          common.ReadJson(filepath.Join(BaseDir, "data/payment.list.json")),
	"get.payment.log.list":      common.ReadJson(filepath.Join(BaseDir, "data/payment.log.list.json")),
	"get.region.son.list":       common.ReadJson(filepath.Join(BaseDir, "data/region.son.list.json")),
	"get.delivery.company.list": common.ReadJson(filepath.Join(BaseDir, "data/delivery.company.list.json")),
	"get.delivery.list":         common.ReadJson(filepath.Join(BaseDir, "data/delivery.list.json")),
	"get.app.list":              common.ReadJson(filepath.Join(BaseDir, "data/app.list.json")),
	"get.app.install.list":      common.ReadJson(filepath.Join(BaseDir, "data/app.install.list.json")),
}

func Default(result string, context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"data":    nil,
		"message": "error",
		"status":  http.StatusNotFound,
	})
}

func doHandle(result string, context *gin.Context) {
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
}

func CommonHandle(context *gin.Context) {
	var defaultParams request.DefaultParams
	if !BindCheck(&defaultParams, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
		return
	}
	rs, ok := MaptoMethod[defaultParams.Method]
	if ok {
		doHandle(rs, context)
	} else {
		Default(rs, context)
	}
}
