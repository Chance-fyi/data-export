package database

import (
	"data-export/pkg/console"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"unsafe"
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

func ScanRows2map(rows *sql.Rows) []map[string]string {
	res := make([]map[string]string, 0)               //  定义结果 map
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址

	}
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		record := make(map[string]string)
		for i, colType := range colTypes {

			if rowValue[i] == nil {
				record[colType.Name()] = ""
			} else {
				record[colType.Name()] = byte2Str(rowValue[i].([]byte))
			}
		}
		res = append(res, record)
	}
	return res
}

func byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
