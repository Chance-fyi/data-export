package model

import "time"

type Role struct {
	Id         uint
	Name       string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
