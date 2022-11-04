package model

import "time"

type Menu struct {
	Id         uint
	Name       string
	Path       string
	ParentId   uint
	CreateTime time.Time `gorm:"autoCreateTime"`
}
