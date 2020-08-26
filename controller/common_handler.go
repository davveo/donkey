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
	"defalut":                     "",
	"get.setting.list":            common.ReadJson(filepath.Join(BaseDir, "data/setting.list.json")),
	"get.upload.module":           common.ReadJson(filepath.Join(BaseDir, "data/upload.module.json")),
	"get.action.log.list":         common.ReadJson(filepath.Join(BaseDir, "data/action.log.list.json")),
	"get.help.router":             common.ReadJson(filepath.Join(BaseDir, "data/help.router.json")),
	"get.auth.rule.list":          common.ReadJson(filepath.Join(BaseDir, "data/auth.rule.list.json")),
	"get.auth.group.list":         common.ReadJson(filepath.Join(BaseDir, "data/auth.group.list.json")),
	"query.app.install.updated":   common.ReadJson(filepath.Join(BaseDir, "data/app.install.json")),
	"get.menu.auth.list":          common.ReadJson(filepath.Join(BaseDir, "data/get.menu.auth.list.json")),
	"get.menu.list":               common.ReadJson(filepath.Join(BaseDir, "data/menu.list.json")),
	"add.admin.item":              common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json")),
	"check.admin.nickname":        common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json")),
	"login.admin.user":            common.ReadJson(filepath.Join(BaseDir, "data/login.admin.user.json")),
	"get.admin.list":              common.ReadJson(filepath.Join(BaseDir, "data/admin.list.json")),
	"get.message.user.unread":     common.ReadJson(filepath.Join(BaseDir, "data/message.user.unread.json")),
	"get.navigation.list":         common.ReadJson(filepath.Join(BaseDir, "data/navigation.list.json")),
	"get.notice.tpl.list":         common.ReadJson(filepath.Join(BaseDir, "data/notice.tpl.list.json")),
	"get.payment.list":            common.ReadJson(filepath.Join(BaseDir, "data/payment.list.json")),
	"get.payment.log.list":        common.ReadJson(filepath.Join(BaseDir, "data/payment.log.list.json")),
	"get.region.son.list":         common.ReadJson(filepath.Join(BaseDir, "data/region.son.list.json")),
	"get.delivery.company.list":   common.ReadJson(filepath.Join(BaseDir, "data/delivery.company.list.json")),
	"get.delivery.list":           common.ReadJson(filepath.Join(BaseDir, "data/delivery.list.json")),
	"get.app.list":                common.ReadJson(filepath.Join(BaseDir, "data/app.list.json")),
	"get.app.install.list":        common.ReadJson(filepath.Join(BaseDir, "data/app.install.list.json")),
	"logout.admin.user":           common.ReadJson(filepath.Join(BaseDir, "data/logout.admin.user.json")),
	"get.app.captcha":             common.ReadJson(filepath.Join(BaseDir, "data/app.captcha.json")),
	"get.brand.select":            common.ReadJson(filepath.Join(BaseDir, "data/brand.select.json")),
	"get.goods.type.select":       common.ReadJson(filepath.Join(BaseDir, "data/goods.type.select.json")),
	"get.goods.admin.list":        common.ReadJson(filepath.Join(BaseDir, "data/goods.admin.list.json")),
	"get.goods.category.list":     common.ReadJson(filepath.Join(BaseDir, "data/goods.category.list.json")),
	"get.goods.spec.page":         common.ReadJson(filepath.Join(BaseDir, "data/goods.spec.page.json")),
	"get.goods.attribute.page":    common.ReadJson(filepath.Join(BaseDir, "data/goods.attribute.page.json")),
	"get.goods.comment.list":      common.ReadJson(filepath.Join(BaseDir, "data/goods.comment.list.json")),
	"get.goods.consult.list":      common.ReadJson(filepath.Join(BaseDir, "data/goods.consult.list.json")),
	"get.order.status.total":      common.ReadJson(filepath.Join(BaseDir, "data/order.status.total.json")),
	"get.order.list":              common.ReadJson(filepath.Join(BaseDir, "data/order.list.json")),
	"get.delivery.select":         common.ReadJson(filepath.Join(BaseDir, "data/delivery.select.json")),
	"get.delivery.company.select": common.ReadJson(filepath.Join(BaseDir, "data/delivery.company.select.json")),
	"get.transaction.list":        common.ReadJson(filepath.Join(BaseDir, "data/transaction.list.json")),
	"get.withdraw.list":           common.ReadJson(filepath.Join(BaseDir, "data/withdraw.list.json")),
	"get.ask.list":                common.ReadJson(filepath.Join(BaseDir, "data/ask.list.json")),
	"get.discount.list":           common.ReadJson(filepath.Join(BaseDir, "data/discount.list.json")),
	"get.promotion.list":          common.ReadJson(filepath.Join(BaseDir, "data/promotion.list.json")),
	"get.coupon.list":             common.ReadJson(filepath.Join(BaseDir, "data/coupon.list.json")),
	"get.user.level.list":         common.ReadJson(filepath.Join(BaseDir, "data/user.level.list.json")),
	"get.coupon.give.list":        common.ReadJson(filepath.Join(BaseDir, "data/coupon.give.list.json")),
	"get.card.list":               common.ReadJson(filepath.Join(BaseDir, "data/card.list.json")),
	"get.card.use.list":           common.ReadJson(filepath.Join(BaseDir, "data/card.use.list.json")),
	"get.user.list":               common.ReadJson(filepath.Join(BaseDir, "data/user.list.json")),
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
