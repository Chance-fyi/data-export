package model

import "time"

type Database struct {
	Id         uint
	Name       string
	Hostname   string
	Port       string
	Username   string
	Password   string
	Database   string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
