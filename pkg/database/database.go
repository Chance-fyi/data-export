package database

import (
	"data-export/pkg/console"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
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

func (db *db) SetConnections(name string, Db *gorm.DB) {
	db.connections[name] = Db
}

func ScanRows2map(rows *sql.Rows) (columns []string, list []map[string]string) {
	columns, _ = rows.Columns()
	columnLength := len(columns)

	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}

	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]string, columnLength)
		for i, data := range cache {
			v := *data.(*interface{})
			switch v.(type) {
			case time.Time:
				item[columns[i]] = v.(time.Time).Format("2006-01-02 15:04:05")
			case []uint8:
				item[columns[i]] = string(v.([]byte))
			}
		}

		list = append(list, item)
	}

	return
}
