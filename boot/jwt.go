package boot

import (
	"data-export/pkg/config"
	"data-export/pkg/jwt"
)

func initJwt() {
	config.UnmarshalKey("jwt", &jwt.Config)
}
