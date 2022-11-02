package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
	"golang.org/x/crypto/bcrypt"
)

func Login(r api.LoginRequest) (model.User, error) {
	var user model.User
	err := g.DB().Where("username = ?", r.Username).First(&user).Error
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))

	return user, err
}
