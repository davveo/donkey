package admin

import (
	"fmt"

	"github.com/davveo/donkey/models"
	"github.com/davveo/donkey/utils/common"
	pass "github.com/davveo/donkey/utils/password"
	"github.com/jinzhu/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Close() {}

func (s *Service) Login(username, password string, lastLogin int64, lastIp, platform string, isGetToken bool) (bool, error) {
	// 根据账号获取
	adminAccount := new(models.AdminAccount)
	accountInstance, err := adminAccount.FindByUserName(s.db, username)
	if err != nil {
		return false, common.ErrAccountNotFound
	}
	if accountInstance.Status == common.NO {
		return false, common.AccountForbbinden
	}
	hashPassword, err := pass.HashPassword(password)
	if err != nil {
		return false, err
	}
	if string(hashPassword) != accountInstance.Password {
		return false, common.ErrAccountOrPassword
	}
	err = adminAccount.Update(s.db, accountInstance, lastLogin, lastIp)
	if err != nil {
		return false, err
	}

	////////////// 返回token信息

	// 如果不需要返回token, 直接返回用户信息
	if !isGetToken {
		return true, nil
	}
	token := new(models.Token)
	tokenInstance, err := token.SetToken(s.db, accountInstance.AdminID,
		accountInstance.GroupID, 1, accountInstance.Username, platform)
	if err != nil {
		return false, nil
	}
	fmt.Println(tokenInstance, accountInstance)
	return true, nil
}

func (s *Service) Logout() bool {
	return true
}
