package boot

import (
	"data-export/pkg/config"
	"data-export/pkg/sqlparse"
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
	//初始化配置
	config.Init()
	initDatabase()
	initRedis()
	initLogger()
	initValidator()

	initJwt()
	initCasbin()
	sqlparse.New()
	initConsumer()
}
