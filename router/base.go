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
		ApiGroup.POST("payment_log", controller.CommonHandle)
		ApiGroup.POST("region", controller.CommonHandle)
		ApiGroup.POST("app", controller.CommonHandle)
		ApiGroup.POST("app_install", controller.CommonHandle)
		ApiGroup.POST("brand", controller.CommonHandle)
		ApiGroup.POST("goods_type", controller.CommonHandle)
		ApiGroup.POST("goods_category", controller.CommonHandle)
		ApiGroup.POST("goods", controller.CommonHandle)
		ApiGroup.POST("spec", controller.CommonHandle)
		ApiGroup.POST("goods_attribute", controller.CommonHandle)
		ApiGroup.POST("goods_comment", controller.CommonHandle)
		ApiGroup.POST("goods_consult", controller.CommonHandle)
		ApiGroup.POST("payment", controller.CommonHandle)
		ApiGroup.POST("order", controller.CommonHandle)
		ApiGroup.POST("delivery", controller.CommonHandle)
		ApiGroup.POST("delivery_item", controller.CommonHandle)
		ApiGroup.POST("transaction", controller.CommonHandle)
		ApiGroup.POST("withdraw", controller.CommonHandle)
		ApiGroup.POST("ask", controller.CommonHandle)
		ApiGroup.POST("discount", controller.CommonHandle)
		ApiGroup.POST("promotion", controller.CommonHandle)
		ApiGroup.POST("coupon", controller.CommonHandle)
		ApiGroup.POST("user_level", controller.CommonHandle)
		ApiGroup.POST("coupon_give", controller.CommonHandle)
		ApiGroup.POST("card", controller.CommonHandle)
		ApiGroup.POST("card_use", controller.CommonHandle)
		ApiGroup.POST("user", controller.CommonHandle)
		ApiGroup.POST("ads_position", controller.CommonHandle)
		ApiGroup.POST("ads", controller.CommonHandle)
		ApiGroup.POST("article_cat", controller.CommonHandle)
		ApiGroup.POST("article", controller.CommonHandle)
		ApiGroup.POST("topic", controller.CommonHandle)
		ApiGroup.POST("storage", controller.CommonHandle)
		ApiGroup.POST("storage_style", controller.CommonHandle)
		ApiGroup.POST("support", controller.CommonHandle)
		ApiGroup.POST("delivery_dist", controller.CommonHandle)
		ApiGroup.POST("friend_link", controller.CommonHandle)
		ApiGroup.POST("qrcode", controller.CommonHandle)
		ApiGroup.POST("admin_login", controller.LoginAdmin)
	}

	return router
}
