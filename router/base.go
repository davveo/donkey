package router

import (
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
				"get.message.user.unread": controller.MessageUnread,
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
				"get.message.user.unread": controller.MessageUnread,
				"add.admin.item":          controller.AddAdminItem,
				"check.admin.nickname":    controller.CheckAdminNickname,
				"login.admin.user":        controller.CheckAdminNickname,
				"get.menu.auth.list":      controller.CheckAdminNickname,
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
				"query.app.install.updated": controller.CheckAdminNickname,
			}

			MaptoMethod[defaultParams.Method](context)
		})

	}

	return router
}
