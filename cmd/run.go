package cmd

import (
	"strconv"
	"time"

	"github.com/davveo/donkey/config"
	"github.com/jinzhu/gorm"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/davveo/donkey/router"
	"github.com/davveo/donkey/services"
	"github.com/davveo/donkey/utils/log"
)

func RunServer(configFile string) error {
	var (
		cfg *config.Config
		db  *gorm.DB
		err error
	)

	// 初始化配置
	if cfg, db, err = initConfig(configFile); err != nil {
		return err
	}
	defer db.Close()

	// 初始化插件 db redis mq
	if err = initPlugin(cfg); err != nil {
		return err
	}

	// 初始化服务
	if err := services.Init(cfg, db); err != nil {
		return err
	}
	defer services.Close()

	// 启动服务
	log.INFO.Println("Starting donkey server on port ", cfg.ServerPort)
	graceful.Run(":"+strconv.Itoa(cfg.ServerPort), 5*time.Second, router.SetupRouter())
	return nil
}
