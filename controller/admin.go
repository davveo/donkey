package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageUnread(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data": map[string]interface{}{
			"admin": map[string]interface{}{
				"username":    "admin55",
				"group_id":    1,
				"nickname":    "admin55",
				"head_pic":    "",
				"last_login":  "2020-08-25 17:10:26",
				"last_ip":     "111.203.244.2",
				"status":      1,
				"create_time": "2020-08-25 14:57:36",
				"update_time": "2020-08-25 17:10:26",
			},
			"token": map[string]interface{}{
				"group_id":        1,
				"token":           "ebb0bdfd1ae40d0424939917835bd635",
				"token_expires":   1600938937,
				"refresh":         "6921718c73ab37b268b3b2a1ff694c93",
				"refresh_expires": 1601025337,
			},
		},
	})
}

func AddAdminItem(context *gin.Context) {

}

func CheckAdminNickname(context *gin.Context) {

}
