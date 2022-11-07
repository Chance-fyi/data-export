package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/casbin"
	"data-export/pkg/g"
	"strconv"
)

func CreateMenu(r api.CreateMenuRequest) error {
	menu := model.Menu{
		Name:     r.Name,
		Path:     r.Path,
		Icon:     r.Icon,
		ParentId: r.ParentId,
	}
	err := g.DB().Create(&menu).Error

	return err
}

func EditMenu(r api.EditMenuRequest) error {
	menu := model.Menu{
		Id:       r.Id,
		Name:     r.Name,
		Path:     r.Path,
		Icon:     r.Icon,
		ParentId: r.ParentId,
	}
	err := g.DB().Model(&menu).Updates(menu).Error

	return err
}

func GetMenu(id uint) (menu api.GetMenuResponse) {
	g.DB().Model(&model.Menu{}).First(&menu, id)
	return
}

func MenuList(r api.MenuListRequest) (menus []api.MenuListItem, count int64) {
	Db := g.DB().Model(&model.Menu{})

	if r.Name != "" {
		Db.Where("name like ?", "%"+r.Name+"%")
	}

	Db.Order("id DESC")
	Db.Count(&count)
	Db.Offset((r.Current - 1) * r.PageSize).Limit(r.PageSize).Find(&menus)

	return
}

func MenuSelectTree(menus []model.Menu, parentId uint) (tree []api.MenuSelectTreeResponse) {
	if menus == nil {
		g.DB().Find(&menus)
	}

	for _, menu := range menus {
		if parentId == menu.ParentId {
			tree = append(tree, api.MenuSelectTreeResponse{
				Value:    menu.Id,
				Title:    menu.Name,
				Children: MenuSelectTree(menus, menu.Id),
			})
		}
	}

	return
}

func UsesMenuList(menus []model.Menu, parentId uint, userId string) (tree []api.UsesMenuList) {
	if menus == nil {
		var menuIds []int
		user, err := g.Casbin().GetRolesForUser(userId)
		if err != nil {
			return nil
		}
		for _, u := range user {
			policy := g.Casbin().GetFilteredPolicy(0, u, casbin.Menu)
			for _, p := range policy {
				menuId, _ := strconv.Atoi(p[2])
				menuIds = append(menuIds, menuId)
			}
		}

		g.DB().Where("id in ?", menuIds).Find(&menus)
	}

	for _, menu := range menus {
		if parentId == menu.ParentId {
			tree = append(tree, api.UsesMenuList{
				Name:   menu.Name,
				Path:   menu.Path,
				Icon:   menu.Icon,
				Routes: UsesMenuList(menus, menu.Id, userId),
			})
		}
	}

	return
}
