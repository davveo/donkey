package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BindCheck(obj interface{}, context *gin.Context) bool {
	if err := context.ShouldBindJSON(&obj); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
