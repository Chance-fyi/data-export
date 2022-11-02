package boot

import (
	"data-export/pkg/config"
	log "data-export/pkg/logger"
)

func initLogger() {
	var cfg log.Config
	config.UnmarshalKey("logger", &cfg)
	log.InitLogger(&cfg)
}
