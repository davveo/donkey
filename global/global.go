package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	GD *gorm.DB
	GR *redis.Client
)
