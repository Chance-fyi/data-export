package model

type Database struct {
	Id       uint
	Name     string
	Hostname string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
	TimeModel
}

func (Database) TableName() string {
	return prefix + "database"
}
