package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateRole(ctx *gin.Context) {
	var r api.CreateRoleRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.CreateRole(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}

func RoleList(ctx *gin.Context) {
	var r api.RoleListRequest
	_ = ctx.ShouldBind(&r)

	roles, count := service.RoleList(r)

	response.Success(ctx, "", api.RoleListResponse{
		Total: count,
		Data:  roles,
	})
}

func GetRole(ctx *gin.Context) {
	var r api.GetRoleRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	role := service.GetRole(r.Id)
	response.Success(ctx, "", role)
}

func EditRole(ctx *gin.Context) {
	var r api.EditRoleRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	err = service.EditRole(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "修改成功")
}
