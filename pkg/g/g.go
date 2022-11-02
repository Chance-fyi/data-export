package g

import (
	"data-export/pkg/database"
	r "data-export/pkg/redis"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

func DB(name ...string) *gorm.DB {
	return database.DB.Connect(name...)
}

func Redis() *redis.Client {
	return r.Redis
}
