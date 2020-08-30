package models

import "github.com/jinzhu/gorm"

type AuthGroup struct {
	gorm.Model
	Name        string `gorm:"column:name;NOT NULL" json:"name"`               // 名称
	Description string `gorm:"column:description;NOT NULL" json:"-"`           // 描述
	Module      string `gorm:"column:module;NOT NULL" json:"-"`                // 所属模块
	System      int    `gorm:"column:system;default:0;NOT NULL" json:"-"`      // 系统保留
	Sort        int    `gorm:"column:sort;default:50;NOT NULL" json:"-"`       // 排序
	Status      int    `gorm:"column:status;default:1;NOT NULL" json:"status"` // 0=禁用 1=启用
}

func (m *AuthGroup) TableName() string {
	return "donkey_auth_group"
}
