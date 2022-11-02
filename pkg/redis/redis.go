package redis

import (
	"github.com/go-redis/redis/v9"
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

var Redis *redis.Client

func NewClient(address string, password string, db int) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}
