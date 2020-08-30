package models

import (
	"github.com/davveo/donkey/utils/common"
	"github.com/jinzhu/gorm"
)

type App struct {
	AppId     int    `gorm:"column:app_id;NOT NULL" json:"app_id"`
	AppName   string `gorm:"column:app_name;NOT NULL" json:"app_name"`             // 名称
	AppKey    int    `gorm:"column:app_key;NOT NULL" json:"app_key"`               // 钥匙
	AppSecret string `gorm:"column:app_secret;NOT NULL" json:"app_secret"`         // 密钥
	Captcha   int    `gorm:"column:captcha;default:0;NOT NULL" json:"captcha"`     // 启用验证码 0=否 1=是
	Status    int    `gorm:"column:status;default:1;NOT NULL" json:"status"`       // 0=禁用 1=启用
	IsDelete  int    `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"` // 0=未删 1=已删
}

func (m *App) TableName() string {
	return "donkey_app"
}

func (m *App) FindByKey(db *gorm.DB, appKey string) (*App, error) {
	t := App{}
	has := db.Where("app_key = ?", appKey).
		Take(&t).RecordNotFound()
	if !has {
		return nil, common.ErrNotFound
	}
	return &t, nil
}

func (m *App) GetAppCaptcha(db *gorm.DB, clientType, appKey string, session bool) map[string]interface{} {
	resp := map[string]interface{}{"captcha": true, "session_id": ""}
	appResult, _ := m.FindByKey(db, appKey)
	if appResult != nil && appResult.Captcha == common.NO {
		resp["captcha"] = false
		return resp
	}

	if session {
		//captcha := captcha.NewCaptcha()
		// -1:游客 0:顾客 1:管理组
		//if clientType == "-1" {
		//	resp["session_id"] = captcha.GetKey(randomstr.GenRandomString(32))
		//} else {
		//	resp["session_id"] = captcha.GetKey(controller.GetClientToken(nil))
		//}
	}
	return resp
}
