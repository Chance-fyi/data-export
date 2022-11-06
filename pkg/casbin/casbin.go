package casbin

import (
	"data-export/pkg/console"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Casbin *casbin.Enforcer

const (
	Menu = "menu"
)

func New(db *gorm.DB, modelStr string) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		console.ExitIf(err)
	}
	m, err := model.NewModelFromString(modelStr)
	if err != nil {
		console.ExitIf(err)
	}
	Casbin, err = casbin.NewEnforcer(m, a)
	if err != nil {
		console.ExitIf(err)
	}
}
