package admin

import (
	"github.com/davveo/donkey/models/response"

	pass "github.com/davveo/donkey/utils/password"

	"github.com/davveo/donkey/models"
	"github.com/davveo/donkey/utils/common"
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

func (s *Service) Login(username, password string, lastLogin int64, lastIp, platform string, isGetToken bool) (*response.AdminResp, error) {
	// 根据账号获取
	adminAccount := new(models.AdminAccount)
	accountInstance, err := adminAccount.FindByUserName(s.db, username)
	if err != nil {
		return nil, common.ErrAccountNotFound
	}
	if accountInstance.Status == common.NO {
		return nil, common.AccountForbbinden
	}

	if !pass.ComparePasswords(accountInstance.Password, []byte(password)) {
		return nil, common.ErrAccountOrPassword
	}
	err = adminAccount.Update(s.db, accountInstance, lastLogin, lastIp)
	if err != nil {
		return nil, err
	}

	////////////// 返回token信息

	// 如果不需要返回token, 直接返回用户信息
	if !isGetToken {
		return nil, nil
	}
	token := &models.Token{}
	tokenInstance, err := token.SetToken(s.db, accountInstance.AdminID,
		accountInstance.GroupID, 1, accountInstance.Username, platform)
	if err != nil {
		return nil, nil
	}

	return &response.AdminResp{
		Admin: accountInstance,
		Token: tokenInstance,
	}, nil
}

func (s *Service) Logout() bool {
	return true
}
