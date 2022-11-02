package boot

import (
	"context"
	"data-export/pkg/config"
	"data-export/pkg/console"
	"data-export/pkg/g"
	"data-export/pkg/redis"
	"fmt"
)

func initRedis() {
	var r redis.Config
	config.UnmarshalKey("redis", &r)
	redis.NewClient(fmt.Sprintf("%v:%v", r.Host, r.Port), r.Password, r.DB)
	if err := g.Redis().Ping(context.Background()).Err(); err != nil {
		console.ExitIf(err)
	}
}
