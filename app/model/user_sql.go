package model

import "time"

type UserSql struct {
	Id         uint
	UserId     uint
	SqlId      uint
	Name       string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
