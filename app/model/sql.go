package model

type Sql struct {
	Id         uint
	Name       string
	Sql        string
	Fields     string
	DatabaseId uint
	TimeModel
}

func (Sql) TableName() string {
	return prefix + "sql"
}
