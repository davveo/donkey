package controller

import (
	"github.com/davveo/donkey/models/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Default(context *gin.Context)  {
	context.JSON(http.StatusNotFound, gin.H{
		"data": nil,
		"message":"error",
		"status":http.StatusNotFound,
	})
}

func CommonHandle(defaultParams request.DefaultParams, context *gin.Context)  {
	MaptoMethod := map[string]gin.HandlerFunc{
		"defalut": Default,
		"get.setting.list": Setting,
		"get.upload.module": Upload,
		"get.action.log.list": ActionLogList,
		"get.help.router": Help,
		"get.auth.rule.list": AuthRuleList,
		"get.auth.group.list": AuthGroupList,
		"query.app.install.updated": AppInstallUpdated,
		"get.menu.auth.list": MenuAuthList,
		"get.menu.list": MenuList,
		"add.admin.item":       AdminItem,
		"check.admin.nickname": CheckAdmin,
		"login.admin.user":     LoginAdmin,
		"get.admin.list":     AdminList,
		"get.message.user.unread": MessageUnRead,
	}
	handlerFunc, ok := MaptoMethod[defaultParams.Method]
	if ok {
		handlerFunc(context)
	} else {
		Default(context)
	}
}
