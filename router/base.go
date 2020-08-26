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
			controller.CommonHandle(defaultParams, context)

		})
		ApiGroup.POST("admin", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("menu", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("app_install.html", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("auth_group", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("auth_rule", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("help", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("action_log", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("setting", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
		ApiGroup.POST("upload", func(context *gin.Context) {
			var defaultParams request.DefaultParams
			if !controller.BindCheck(&defaultParams, context) {
				response.FailWithMessage(response.ParamValidateFailed, context)
				return
			}
			controller.CommonHandle(defaultParams, context)
		})
	}

	return router
}
