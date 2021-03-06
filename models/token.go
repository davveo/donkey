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
	TokenID uint32 `json:"-"`
	// ClientID 编号
	ClientID uint `json:"-"`
	// GroupID 用户组Id
	GroupID uint `json:"group_id"`
	// Username 账号
	Username string `json:"-"`
	// ClientType 0=顾客 1=管理组
	ClientType uint32 `json:"-"`
	// Platform 来源终端
	Platform string `json:"-"`
	// Code 随机密钥
	Code string `json:"-"`
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

func (token *Token) Create(db *gorm.DB, instance *Token) error {
	return db.Create(instance).Error
}

func (token *Token) Update(db *gorm.DB, instance, UpdateInstance *Token) error {
	return db.Model(&instance).Updates(&UpdateInstance).Error
}

func (token *Token) FindByClientId(db *gorm.DB, clientId uint, clientType uint32, platform string) (*Token, error) {
	t := Token{}
	has := db.Where("client_id = ?", clientId).
		Where("client_type = ?", clientType).
		Where("platform = ?", platform).
		Take(&t).RecordNotFound()
	if !has {
		return nil, common.ErrNotFound
	}
	return &t, nil
}

func (token *Token) FindByToken(db *gorm.DB, tokenStr string) (*Token, error) {
	t := Token{}
	has := db.Where("token = ?", tokenStr).
		Take(&t).RecordNotFound()
	if !has {
		return nil, common.ErrNotFound
	}
	return &t, nil
}

func (token *Token) FindByClientIdAndToken(db *gorm.DB, clientId, clientType uint32, oldToken string) (*Token, error) {
	t := Token{}
	has := db.Where("client_id = ?", clientId).
		Where("client_type = ?", clientType).
		Where("token = ?", oldToken).
		Take(&t).RecordNotFound()
	if !has {
		return nil, common.ErrNotFound
	}
	return &t, nil
}

func (token *Token) SetToken(db *gorm.DB, clientId, groupId uint, clientType uint32, username, platform string) (*Token, error) {
	code := randomstr.GenRandomString(32)
	tokenStr := decry.UserMd5(fmt.Sprintf("%d%d%s", clientId, clientType, code))
	refreshStr := decry.UserMd5(fmt.Sprintf("%s%s", randomstr.GenRandomString(32), tokenStr))
	expires := time.Now().Unix() + (30 * 24 * 60 * 60) // 30天

	instance, err := token.FindByClientId(db, clientId, clientType, platform)
	if err != nil && instance == nil {
		// 创建操作
		t := &Token{
			ClientID:       clientId,
			GroupID:        groupId,
			Username:       username,
			ClientType:     clientType,
			Platform:       platform,
			Code:           code,
			Token:          tokenStr,
			TokenExpires:   expires,
			Refresh:        refreshStr,
			RefreshExpires: expires + (1 * 24 * 60 * 60),
		}
		if err = token.Create(db, t); err != nil {
			return nil, err
		}
		return t, nil
	} else {
		// 更新操作
		t := &Token{
			Username:       username,
			Code:           code,
			Token:          tokenStr,
			TokenExpires:   expires,
			Refresh:        refreshStr,
			RefreshExpires: expires + (1 * 24 * 60 * 60),
		}
		// 执行更新操作
		err = token.Update(db, instance, t)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}

// 刷新token
// clientType  0/1 顾客或者管理组
// refresh  刷新令牌
// oldToken 原授权令牌
func (token *Token) RefreshToken(db *gorm.DB, clientID, clientType uint32, refresh, oldToken string) (*Token, error) {

	instance, err := token.FindByClientIdAndToken(db, clientID, clientType, oldToken)
	if err != nil {
		return nil, common.ErrRefreshNotFound
	}
	now := time.Now().Unix()
	if now > instance.RefreshExpires {
		return nil, common.ErrRefreshInValid
	}
	if instance.Refresh != refresh {
		return nil, common.ErrRefresh
	}
	// 准备更新数据
	code := randomstr.GenRandomString(32)
	tokenStr := decry.UserMd5(fmt.Sprintf("%d%d%s", clientID, clientType, code))
	expires := time.Now().Unix() + (30 * 24 * 60 * 60) // 30天
	refreshStr := decry.UserMd5(fmt.Sprintf("%s%s", randomstr.GenRandomString(32), tokenStr))
	t := &Token{
		Code:           code,
		Token:          tokenStr,
		TokenExpires:   expires,
		Refresh:        refreshStr,
		RefreshExpires: expires + (1 * 24 * 60 * 60),
	}
	// 执行更新操作
	if err = token.Update(db, instance, t); err != nil {
		return nil, err
	}
	return t, nil
}
