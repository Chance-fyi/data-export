package service

import (
	"data-export/app/model"
	"data-export/pkg/g"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user model.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)

	err = g.DB().Create(&user).Error

	return err
}
