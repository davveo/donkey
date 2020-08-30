package models

import (
	"github.com/davveo/donkey/models/request"
	"github.com/davveo/donkey/utils/common"
	"github.com/jinzhu/gorm"
	"time"
)

//type Admin struct {
//	gorm.Model
//	// Username 账号
//	Username string `json:"username"`
//	// Password 密码
//	Password string `json:"-"`
//	// GroupID 用户组Id
//	GroupID uint32 `json:"group_id"`
//	// Nickname 昵称
//	Nickname string `json:"nickname"`
//	// HeadPic 头像
//	HeadPic string `json:"head_pic"`
//	// LastLogin 最后登录日期
//	LastLogin int64 `json:"last_login"`
//	// LastIP 最后登录ip
//	LastIP string `json:"last_ip"`
//	// Status 0=禁用 1=启用
//	Status int8 `json:"status"`
//	// IsDelete 0=未删 1=已删
//	IsDelete  int8 `json:"-"`
//	AuthGroup `json:"get_auth_group"`
//}

type Admin struct {
	gorm.Model
	Username  string    `json:"username"`                                 // 账号
	Password  string    `json:"password"`                                 // 密码
	GroupId   uint      `json:"group_id"`                                 // 用户组Id
	Nickname  string    `gorm:"column:nickname;NOT NULL" json:"nickname"` // 昵称
	HeadPic   string    `json:"head_pic"`                                 // 头像
	LastLogin time.Time `json:"last_login"`                               // 最后登录日期
	LastIp    string    `json:"last_ip"`                                  // 最后登录ip
	Status    int       `json:"status"`                                   // 0=禁用 1=启用
	IsDelete  int       `json:"is_delete"`                                // 0=未删 1=已删
}

func (admin *Admin) TableName() string {
	return "donkey_admin"
}

func (admin Admin) FindByUserName(db *gorm.DB, username string) (*Admin, error) {
	adminS := &Admin{}
	notFound := db.Where("username = ?", username).
		Take(&adminS).RecordNotFound()
	if notFound {
		return nil, common.ErrNotFound
	}
	return adminS, nil
}

func (admin Admin) Update(db *gorm.DB, instance *Admin, lastLogin time.Time, lastIp string) error {
	return db.Model(&instance).Updates(Admin{LastLogin: lastLogin, LastIp: lastIp}).Error
}

func (admin Admin) GetAdminList(
	db *gorm.DB, pageInfo request.PageInfo,
	order string, desc bool) (err error, list interface{}, total int) {
	var adminList []Admin
	db = db.Model(&Admin{})
	err = db.Count(&total).Error
	if err != nil {
		return err, adminList, total
	}

	db = db.Limit(pageInfo.PageSize).Offset(pageInfo.Page)
	if order != "" {
		var OrderStr string
		if desc {
			OrderStr = order + "desc"
		} else {
			OrderStr = order
		}
		err = db.Order(OrderStr, true).Find(&adminList).Error
	} else {
		err = db.Preload("AuthGroup").Order("username", true).Find(&adminList).Error
		//err = db.Table("donkey_admin").Order("username", true).Joins("JOIN donkey_auth_group ON donkey_auth_group.group_id = donkey_admin.group_id").Find(&adminList).Error
	}
	return err, adminList, total
}
