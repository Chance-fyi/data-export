package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/casbin"
	"data-export/pkg/g"
	"strconv"
)

func CreateRole(r api.CreateRoleRequest) error {
	role := model.Role{Name: r.Name}
	err := g.DB().Create(&role).Error
	if err != nil {
		return err
	}

	var policies [][]string
	roleId := strconv.Itoa(int(role.Id))
	for _, id := range r.MenuIds {
		policies = append(policies, []string{roleId, casbin.Menu, strconv.Itoa(id)})
	}

	_, err = g.Casbin().AddPolicies(policies)

	return err
}

func RoleList(r api.RoleListRequest) (roles []api.RoleListItem, count int64) {
	Db := g.DB().Model(model.Role{})

	if r.Name != "" {
		Db.Where("name like ?", "%"+r.Name+"%")
	}

	Db.Order("id DESC")
	Db.Count(&count)
	Db.Offset((r.Current - 1) * r.PageSize).Limit(r.PageSize).Find(&roles)

	return
}

func GetRole(id uint) (role api.GetRoleResponse) {
	g.DB().Model(&model.Role{}).First(&role, id)

	policies := g.Casbin().GetFilteredPolicy(0, strconv.Itoa(int(role.Id)))
	for _, policy := range policies {
		if policy[1] != casbin.Menu {
			continue
		}
		menuId, _ := strconv.Atoi(policy[2])
		role.MenuIds = append(role.MenuIds, menuId)
	}
	return
}

func EditRole(r api.EditRoleRequest) error {
	role := model.Role{
		Id:   r.Id,
		Name: r.Name,
	}
	err := g.DB().Model(&role).Updates(role).Error
	if err != nil {
		return err
	}

	roleId := strconv.Itoa(int(role.Id))
	oldPolicies := g.Casbin().GetFilteredPolicy(0, roleId)
	_, err = g.Casbin().RemovePolicies(oldPolicies)
	if err != nil {
		return err
	}

	var policies [][]string
	for _, id := range r.MenuIds {
		policies = append(policies, []string{roleId, casbin.Menu, strconv.Itoa(id)})
	}
	_, err = g.Casbin().AddPolicies(policies)

	return err
}

func UserRoleList() (list []api.UserRoleListResponse) {
	g.DB().Model(model.Role{}).Find(&list)
	return
}
