package services

import (
	"reflect"

	"github.com/davveo/donkey/services/healthcheck"

	"github.com/davveo/donkey/config"
	"github.com/jinzhu/gorm"
)

var (
	HealthService healthcheck.ServiceInterface
)

func UseHealthService(h healthcheck.ServiceInterface) {
	HealthService = h
}

func Init(cfg *config.Config, db *gorm.DB) error {
	if nil == reflect.TypeOf(HealthService) {
		HealthService = healthcheck.NewService(db)
	}

	return nil
}

func Close() {
	HealthService.Close()
}
