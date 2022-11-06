package boot

import (
	"data-export/pkg/casbin"
	"data-export/pkg/config"
	"data-export/pkg/g"
)

func initCasbin() {
	casbin.New(g.DB(), config.GetString("casbin.model"))
}
