package controller

import (
	"net/http"

	"github.com/davveo/donkey/models/request"
	"github.com/davveo/donkey/utils/response"
	"github.com/gin-gonic/gin"
)

var MaptoMethod = map[string]gin.HandlerFunc{
	"defalut":                   Default,
	"get.setting.list":          Setting,
	"get.upload.module":         Upload,
	"get.action.log.list":       ActionLogList,
	"get.help.router":           Help,
	"get.auth.rule.list":        AuthRuleList,
	"get.auth.group.list":       AuthGroupList,
	"query.app.install.updated": AppInstallUpdated,
	"get.menu.auth.list":        MenuAuthList,
	"get.menu.list":             MenuList,
	"add.admin.item":            AdminItem,
	"check.admin.nickname":      CheckAdmin,
	"login.admin.user":          LoginAdmin,
	"get.admin.list":            AdminList,
	"get.message.user.unread":   MessageUnRead,
}

func Default(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{
		"data":    nil,
		"message": "error",
		"status":  http.StatusNotFound,
	})
}

func CommonHandle(context *gin.Context) {
	var defaultParams request.DefaultParams
	if !BindCheck(&defaultParams, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
		return
	}

	handlerFunc, ok := MaptoMethod[defaultParams.Method]
	if ok {
		handlerFunc(context)
	} else {
		Default(context)
	}
}
