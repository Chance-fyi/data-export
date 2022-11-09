package model

import "time"

type Export struct {
	Id         uint
	UserId     uint
	UserSqlId  uint
	SqlId      uint
	DatabaseId uint
	Filename   string
	Sql        string
	Path       string
	Status     uint //0 导出中  1 未下载  2 已下载  3 失败  4 删除
	ErrorMsg   string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
