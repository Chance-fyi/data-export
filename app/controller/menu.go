package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateMenu(ctx *gin.Context) {
	var r api.CreateMenuRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.CreateMenu(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}

func EditMenu(ctx *gin.Context) {
	var r api.EditMenuRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.EditMenu(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "修改成功")
}

func GetMenu(ctx *gin.Context) {
	var r api.GetMenuRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	menu := service.GetMenu(r.Id)
	response.Success(ctx, "", menu)
}

func MenuList(ctx *gin.Context) {
	var r api.MenuListRequest
	_ = ctx.ShouldBind(&r)

	menus, count := service.MenuList(r)

	response.Success(ctx, "", api.MenuListResponse{
		Total: count,
		Data:  menus,
	})
}

func MenuSelectTree(ctx *gin.Context) {
	tree := service.MenuSelectTree(nil, 0)
	response.Success(ctx, "", tree)
}
