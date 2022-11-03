package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
	"github.com/gin-gonic/gin"
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

func UsesMenuList(ctx *gin.Context, menus []model.Menu, parentId uint) (tree []api.UsesMenuList) {
	if menus == nil {
		g.DB().Find(&menus)
	}

	for _, menu := range menus {
		if parentId == menu.ParentId {
			tree = append(tree, api.UsesMenuList{
				Name:   menu.Name,
				Path:   menu.Path,
				Icon:   menu.Icon,
				Routes: UsesMenuList(ctx, menus, menu.Id),
			})
		}
	}

	return
}
