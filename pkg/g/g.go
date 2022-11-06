package g

import (
	c "data-export/pkg/casbin"
	"data-export/pkg/database"
	r "data-export/pkg/redis"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

func DB(name ...string) *gorm.DB {
	return database.DB.Connect(name...)
}

func Redis() *redis.Client {
	return r.Redis
}

func Casbin() *casbin.Enforcer {
	return c.Casbin
}
