package database

import (
	"data-export/pkg/console"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Default     string
	Connections map[string]connection
}

type connection struct {
	Type     string
	Hostname string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
	Prefix   string
}

type db struct {
	defaultConnect string
	connections    map[string]*gorm.DB
}

var DB = db{
	defaultConnect: "",
	connections:    map[string]*gorm.DB{},
}

func (db *db) SetDefaultConnect(name string) {
	db.defaultConnect = name
}

func (db *db) GetDefaultConnect() string {
	return db.defaultConnect
}

func (db *db) Connect(name ...string) *gorm.DB {
	if len(name) > 0 {
		connection, ok := db.connections[name[0]]
		if !ok {
			console.Errorp(fmt.Sprintf("%s:connection does not exist", name[0]))
		}
		return connection
	}
	return db.connections[db.defaultConnect]
}

func (db *db) CreateConnection(name string, dialector gorm.Dialector, cfg connection) {
	open, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix,
			SingularTable: true,
		},
	})
	console.ExitIf(err)
	db.connections[name] = open
}
