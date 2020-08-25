package cmd

import (
	"github.com/davveo/donkey/config"
	database "github.com/davveo/donkey/plugin/db"
	"github.com/jinzhu/gorm"
)

func initConfig(configFile string) (*config.Config, *gorm.DB, error) {
	cfg := config.NewConfig(configFile)

	db, err := database.NewDataBase(cfg)
	if err != nil {
		return nil, nil, err
	}
	return cfg, db, nil
}

func initPlugin(*config.Config) error {
	return nil
}
