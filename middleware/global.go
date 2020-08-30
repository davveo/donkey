package middleware

import (
	"github.com/davveo/donkey/models/request"
	"github.com/gin-gonic/gin"
)

//func checkParams(context *gin.Context, defaultParams request.DefaultParams) bool {
//	return controller.BindCheck(&defaultParams, context)
//}
//
//func checkToken(context *gin.Context, defaultParams request.DefaultParams) bool {
//	// Token为空则表示以游客身份访问
//	if defaultParams.Token == "" {
//		return true
//	}
//	// 从本地数据库获取Token
//	token := models.Token{}
//	byToken, _ := token.FindByToken(nil, defaultParams.Token)
//	if byToken != nil {
//		// 验证token是否过期
//		if byToken.TokenExpires != 0 || time.Now().Unix() > byToken.TokenExpires {
//			return false
//		}
//		// 还原token加密过程
//		tokenStr := decry.UserMd5(
//			fmt.Sprintf("%d%d%s", byToken.ClientID, byToken.ClientType, byToken.Code))
//		if defaultParams.Token != tokenStr {
//			// token不合法
//			return false
//		}
//		context.Set("client", map[string]interface{}{
//			"client_id":   byToken.ClientID,
//			"type":        byToken.ClientType,
//			"client_name": byToken.Username,
//			"group_id":    byToken.GroupID,
//			"token":       defaultParams.Token,
//		})
//		/*
//			设置缓存
//			$cacheTag = 'token:' . (is_client_admin() ? 'admin_' : 'user_') . get_client_id();
//			Cache::tag($cacheTag, ['token:' . $this->token]);
//		*/
//	} else {
//		return false
//	}
//	return true
//}
//func checkAuth(context *gin.Context, defaultParams request.DefaultParams) bool {
//	return false
//}
//func checkApp(context *gin.Context, defaultParams request.DefaultParams) bool {
//	//白名单中排除的接口
//	exclude := []string{"login.admin.user", "login.user.user"}
//	exists, _ := common.InArray(defaultParams.Method, exclude)
//	if exists {
//		// 从app获取信息
//	}
//	return false
//}
func checkSign(context *gin.Context, defaultParams request.DefaultParams) bool {

	return false
}

//func respondWithError(code int, message string, context *gin.Context) {
//	resp := map[string]string{"error": message}
//	context.JSON(code, resp)
//	context.Abort()
//}
//
//func GlobalMiddleWare() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		var (
//			defaultParams request.DefaultParams
//		)
//		// 验证params
//		if checkParams(context, defaultParams) {
//			respondWithError(http.StatusInternalServerError, "参数校验失败", context)
//			return
//		}
//
//		// 验证token
//		if checkToken(context, defaultParams) {
//			// 未授权，请重新登录(401)
//			respondWithError(http.StatusUnauthorized, "未授权，请重新登录", context)
//			return
//		}
//
//		// 验证auth
//		if checkAuth(context, defaultParams) {
//			// 拒绝访问(403)
//			respondWithError(http.StatusForbidden, "", context)
//			return
//		}
//
//		// 验证app
//		if checkApp() {
//			respondWithError(http.StatusInternalServerError, "", context)
//			return
//		}
//		// 验证Sign
//		if checkSign() {
//			respondWithError(http.StatusInternalServerError, "", context)
//			return
//		}
//
//		context.Set("client23", map[string]interface{}{
//			"name": "chenwei",
//			"age":  187,
//		})
//		context.Next()
//	}
//}

func GlobalMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
