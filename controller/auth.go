package controller

import (
	"encoding/json"
	"fmt"
	"github.com/davveo/donkey/fake"
	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func AuthGroupList(context *gin.Context)  {
	var authGroupList fake.AuthGroupList
	result := common.ReadJson(filepath.Join(BaseDir, "data/auth.group.list.json"))
	err := json.Unmarshal([]byte(result), &authGroupList)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  authGroupList.Status,
		"message": authGroupList.Message,
		"data":    authGroupList.Data,
	})
}