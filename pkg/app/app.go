package app

import (
	"data-export/app/model"
	"data-export/pkg/config"
	"data-export/pkg/console"
	"github.com/gin-gonic/gin"
	"time"
)

// TimeNowInTimezone get current time, support setting time zone
func TimeNowInTimezone() time.Time {
	location, err := time.LoadLocation(config.GetString("app.timezone"))
	console.ExitIf(err)
	return time.Now().In(location)
}

func IsDebug() bool {
	return config.GetBool("app.debug")
}

func Name() string {
	return config.GetString("app.name")
}

func JwtUser(ctx *gin.Context) (user model.User) {
	value, _ := ctx.Get("jwtUser")
	user = value.(model.User)
	return
}
