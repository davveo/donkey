package models

type AuthRule struct {
	RuleId   int    `gorm:"column:rule_id;NOT NULL" json:"rule_id"`
	Module   string `gorm:"column:module;NOT NULL" json:"module"`           // 所属模块
	GroupId  int    `gorm:"column:group_id;NOT NULL" json:"group_id"`       // 用户组Id
	Name     string `gorm:"column:name;NOT NULL" json:"name"`               // 规则名称
	MenuAuth string `gorm:"column:menu_auth;NOT NULL" json:"menu_auth"`     // 菜单权限
	LogAuth  string `gorm:"column:log_auth;NOT NULL" json:"log_auth"`       // 记录权限
	Sort     int    `gorm:"column:sort;default:50;NOT NULL" json:"sort"`    // 排序
	Status   int    `gorm:"column:status;default:1;NOT NULL" json:"status"` // 0=禁用 1=启用
}

func (m *AuthRule) TableName() string {
	return "donkey_auth_rule"
}
