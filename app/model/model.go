package model

import "time"

const prefix = "de_"

type TimeModel struct {
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}
