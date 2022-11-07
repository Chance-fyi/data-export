package model

import "time"

type Menu struct {
	Id         uint
	Name       string
	Path       string
	Icon       string
	ParentId   uint
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
