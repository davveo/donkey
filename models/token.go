package models

import (
	"fmt"
	"time"

	"github.com/davveo/donkey/utils/common"
	"github.com/davveo/donkey/utils/decry"
	"github.com/davveo/donkey/utils/randomstr"
	"github.com/jinzhu/gorm"
)

type Token struct {
	TokenID uint32 `json:"token_id"`
	// ClientID 编号
	ClientID uint32 `json:"client_id"`
	// GroupID 用户组Id
	GroupID uint32 `json:"group_id"`
	// Username 账号
	Username string `json:"username"`
	// ClientType 0=顾客 1=管理组
	ClientType uint8 `json:"client_type"`
	// Platform 来源终端
	Platform string `json:"platform"`
	// Code 随机密钥
	Code string `json:"code"`
	// Token 授权码
	Token string `json:"token"`
	// TokenExpires 授权码过期时间
	TokenExpires int64 `json:"token_expires"`
	// Refresh 刷新授权码
	Refresh string `json:"refresh"`
	// RefreshExpires 刷新授权码过期时间
	RefreshExpires int64 `json:"refresh_expires"`
}

func (Token) TableName() string {
	return "donkey_token"
}

func (token *Token) Update(db *gorm.DB, instance *Token, lastLogin int64, lastIp string) error {
	return db.Model(&instance).Updates(Token{LastLogin: lastLogin, LastIP: lastIp}).Error
}

func (token *Token) FindByClientId(db *gorm.DB, clientId uint32, clientType uint8, platform string) (*Token, error) {
	t := new(Token)
	has := db.Where("client_id = ?", clientId).
		Where("client_type = ?", clientType).
		Where("platform = ?", platform).
		Take(&t).RecordNotFound()
	if !has {
		return nil, common.ErrNotFound
	}
	return t, nil
}

func (token *Token) Save(db *gorm.DB, instance *Token) error {
	return db.Create(instance).Error
}

func (token *Token) SetToken(db *gorm.DB, adminId, groupId uint32, clientType uint8, username, platform string) (*Token, error) {
	code := randomstr.GenRandomString(32)
	tokenStr := decry.UserMd5(fmt.Sprintf("%d%d%s", adminId, clientType, code))
	refreshStr := decry.UserMd5(fmt.Sprintf("%s%s", randomstr.GenRandomString(32), tokenStr))
	expires := time.Now().Unix() + (30 * 24 * 60 * 60) // 30天
	// TODO
	//instance, err := token.FindByClientId(db, adminId, clientType, platform)
	//if err != nil {
	//	return nil, err
	//}
	//t := &Token{
	//	ClientID:       adminId,
	//	GroupID:        groupId,
	//	Username:       username,
	//	ClientType:     clientType,
	//	Platform:       platform,
	//	Code:           code,
	//	Token:          tokenStr,
	//	TokenExpires:   expires,
	//	Refresh:        refreshStr,
	//	RefreshExpires: expires + (1 * 24 * 60 * 60),
	//}
	//err = token.Save(db, t)
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println(refreshStr, expires)
	return nil, nil
}
