package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func CreateUser(r api.CreateUserRequest) error {
	user := model.User{
		Username: r.Username,
		Password: r.Password,
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)

	err = g.DB().Create(&user).Error
	if err != nil {
		return err
	}

	userId := strconv.Itoa(int(user.Id))
	for _, id := range r.RoleIds {
		_, err = g.Casbin().AddRoleForUser(userId, id)
		if err != nil {
			return err
		}
	}

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

	for i, user := range users {
		roleIds, _ := g.Casbin().GetRolesForUser(strconv.Itoa(int(user.Id)))
		var roles []model.Role
		g.DB().Where("id in ?", roleIds).Find(&roles)
		for _, role := range roles {
			users[i].Role = append(users[i].Role, role.Name)
		}
	}

	return
}

func GetUser(id uint) (user api.GetUserResponse) {
	g.DB().Model(&model.User{}).First(&user, id)
	roleIds, err := g.Casbin().GetRolesForUser(strconv.Itoa(int(user.Id)))
	if err == nil {
		user.RoleIds = roleIds
	}
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
	if err != nil {
		return err
	}

	userId := strconv.Itoa(int(user.Id))
	_, err = g.Casbin().DeleteRolesForUser(userId)
	if err != nil {
		return err
	}
	for _, id := range r.RoleIds {
		_, err = g.Casbin().AddRoleForUser(userId, id)
		if err != nil {
			return err
		}
	}

	return err
}

func UserSelectList() (list []api.UserSelectListResponse) {
	g.DB().Model(model.User{}).Find(&list)
	return
}
