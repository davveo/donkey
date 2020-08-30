package initialize

import (
	"github.com/davveo/donkey/global"
	"github.com/jinzhu/gorm"
	"os"
)

func Mysql() {
	admin := global.GD.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		os.Exit(0)
	} else {
		global.GD = db
		global.GD.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GD.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GD.LogMode(admin.LogMode)
	}
}
