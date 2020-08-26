package router

import (
	"github.com/davveo/donkey/controller"
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
		ApiGroup.POST("message", controller.CommonHandle)
		ApiGroup.POST("admin", controller.CommonHandle)
		ApiGroup.POST("menu", controller.CommonHandle)
		ApiGroup.POST("app_install.html", controller.CommonHandle)
		ApiGroup.POST("auth_group", controller.CommonHandle)
		ApiGroup.POST("auth_rule", controller.CommonHandle)
		ApiGroup.POST("help", controller.CommonHandle)
		ApiGroup.POST("action_log", controller.CommonHandle)
		ApiGroup.POST("setting", controller.CommonHandle)
		ApiGroup.POST("upload", controller.CommonHandle)
		ApiGroup.POST("navigation", controller.CommonHandle)
		ApiGroup.POST("notice_tpl", controller.CommonHandle)
		ApiGroup.POST("payment", controller.CommonHandle)
		ApiGroup.POST("payment_log", controller.CommonHandle)
		ApiGroup.POST("region", controller.CommonHandle)
		ApiGroup.POST("delivery_item", controller.CommonHandle)
		ApiGroup.POST("delivery", controller.CommonHandle)
		ApiGroup.POST("app", controller.CommonHandle)
		ApiGroup.POST("app_install", controller.CommonHandle)
		ApiGroup.POST("admin_login", controller.LoginAdmin)
	}

	return router
}
