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

func MessageUnRead(context *gin.Context) {
	var messageUnread fake.MessageUnread
	result := common.ReadJson(filepath.Join(BaseDir, "data/message.user.unread.json"))
	err := json.Unmarshal([]byte(result), &messageUnread)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  messageUnread.Status,
		"message": messageUnread.Message,
		"data":    messageUnread.Data,
	})
}
