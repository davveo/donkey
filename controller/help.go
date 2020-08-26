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

func Help(context *gin.Context)  {
	var help fake.Help
	result := common.ReadJson(filepath.Join(BaseDir, "data/auth.group.list.json"))
	err := json.Unmarshal([]byte(result), &help)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  help.Status,
		"message": help.Message,
		"data":    help.Data,
	})
}