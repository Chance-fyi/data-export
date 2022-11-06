package service

import (
	"data-export/app/api"
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

func UserList(r api.UserListRequest) (users []api.UserListItem, count int64) {
	Db := g.DB().Model(model.User{})

	if r.Username != "" {
		Db.Where("username like ?", "%"+r.Username+"%")
	}

	Db.Order("id DESC")
	Db.Count(&count)
	Db.Offset((r.Current - 1) * r.PageSize).Limit(r.PageSize).Find(&users)

	return
}

func GetUser(id uint) (user api.GetUserResponse) {
	g.DB().Model(&model.User{}).First(&user, id)
	return
}

func EditUser(r api.EditUserRequest) error {
	user := model.User{
		Id:       r.Id,
		Username: r.Username,
	}
	if r.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(password)
	}

	err := g.DB().Model(&user).Updates(user).Error

	return err
}
