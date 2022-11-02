package boot

import (
	"data-export/pkg/config"
	"data-export/pkg/database"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	var cfg database.Config
	config.UnmarshalKey("database", &cfg)
	database.DB.SetDefaultConnect(cfg.Default)
	setupDb(cfg)
}

func setupDb(cfg database.Config) {
	var dialector gorm.Dialector
	for name, conn := range cfg.Connections {
		switch conn.Type {
		case "mysql":
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
				conn.Username,
				conn.Password,
				conn.Hostname,
				conn.Port,
				conn.Database,
				conn.Charset,
			)
			dialector = mysql.New(mysql.Config{
				DSN: dsn,
			})
		default:
			panic(errors.New("database connection not supported"))
		}
		database.DB.CreateConnection(name, dialector, conn)
	}
}
