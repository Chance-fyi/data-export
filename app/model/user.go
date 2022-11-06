package model

import "time"

type User struct {
	Id         uint
	Username   string
	Password   string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
