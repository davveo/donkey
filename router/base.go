package router

import (
	"fmt"
	"github.com/davveo/donkey/controller"
	"github.com/davveo/donkey/models/request"
	"github.com/davveo/donkey/utils/response"

	"github.com/davveo/donkey/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 中间件管理
	router.Use(middleware.CorsMiddleWare())
	router.Use(middleware.JwtMiddleWare())

	// 路由管理
	ApiGroup := router.Group("api/v1")
	{
		ApiGroup.POST("message", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"get.message.user.unread": controller.MessageUnRead,
			}

			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("admin", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"add.admin.item":       controller.AdminItem,
				"check.admin.nickname": controller.CheckAdmin,
				"login.admin.user":     controller.LoginAdmin,
				"get.admin.list":     controller.AdminList,
			}
			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("menu", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"get.menu.auth.list": controller.MenuAuthList,
				"get.menu.list": controller.MenuList,
			}

			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("app_install.html", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"query.app.install.updated": controller.AppInstallUpdated,
			}

			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("auth_group", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"get.auth.group.list": controller.AuthGroupList,
			}

			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("auth_rule", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"get.auth.rule.list": controller.AuthRuleList,
			}

			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("help", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"get.help.router": controller.Help,
			}

			MaptoMethod[defaultParams.Method](context)
		})
		ApiGroup.POST("action_log", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			MaptoMethod := map[string]gin.HandlerFunc{
				"get.action.log.list": controller.ActionLogList,
			}

			handlerFunc, ok := MaptoMethod[defaultParams.Method]
			fmt.Println(ok)
			handlerFunc(context)
		})
	}

	return router
}
