package controller

import (
	"encoding/json"
	"fmt"
	"github.com/davveo/donkey/fake"
	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

var (
	BaseDir, _ = os.Getwd()
)

func MenuAuthList(context *gin.Context) {
	var menuList fake.MenuList
	result := common.ReadJson(filepath.Join(BaseDir, "data/get.menu.auth.list.json"))
	err := json.Unmarshal([]byte(result), &menuList)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  menuList.Status,
		"message": menuList.Message,
		"data":    menuList.Data,
	})
}
