package app

import (
	"data-export/pkg/config"
	"data-export/pkg/console"
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
