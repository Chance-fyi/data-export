package model

import "time"

type Sql struct {
	Id         uint
	Sql        string
	Fields     string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

func (Sql) TableName() string {
	return prefix + "sql"
}
