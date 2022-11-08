package model

import (
	"data-export/pkg/g"
	"strconv"
	"time"
)

type User struct {
	Id         uint
	Username   string
	Password   string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

func (u User) IsAdmin() bool {
	role := Role{}
	g.DB().Where("name = ?", "admin").First(&role)
	if role.Id == 0 {
		return false
	}
	b, err := g.Casbin().HasRoleForUser(strconv.Itoa(int(u.Id)), strconv.Itoa(int(role.Id)))
	return err == nil && b
}
