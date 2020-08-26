package models

import (
	"github.com/davveo/donkey/utils/common"
	"github.com/jinzhu/gorm"
)

type AdminAccount struct {
	AdminID uint32 `json:"admin_id"`
	// Username 账号
	Username string `json:"username"`
	// Password 密码
	Password string `json:"password"`
	// GroupID 用户组Id
	GroupID uint32 `json:"group_id"`
	// Nickname 昵称
	Nickname string `json:"nickname"`
	// HeadPic 头像
	HeadPic string `json:"head_pic"`
	// LastLogin 最后登录日期
	LastLogin int64 `json:"last_login"`
	// LastIP 最后登录ip
	LastIP string `json:"last_ip"`
	// Status 0=禁用 1=启用
	Status int8 `json:"status"`
	// IsDelete 0=未删 1=已删
	IsDelete int8 `json:"is_delete"`
	// CreateTime 创建日期
	CreateTime int32 `json:"create_time"`
	// UpdateTime 更新日期
	UpdateTime int32 `json:"update_time"`
}

func (AdminAccount) TableName() string {
	return "donkey_admin"
}

func (adminAccount *AdminAccount) FindByUserName(db *gorm.DB, username string) (*AdminAccount, error) {
	account := new(AdminAccount)
	notFound := db.Where("username = ?", username).
		Take(&account).RecordNotFound()
	if notFound {
		return nil, common.ErrNotFound
	}
	return account, nil
}

func (adminAccount *AdminAccount) Update(db *gorm.DB, instance *AdminAccount, lastLogin int64, lastIp string) error {
	return db.Model(&instance).Updates(AdminAccount{LastLogin: lastLogin, LastIP: lastIp}).Error
}
