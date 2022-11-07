package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var r api.CreateUserRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.CreateUser(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}

func UserList(ctx *gin.Context) {
	var r api.UserListRequest
	_ = ctx.ShouldBind(&r)

	users, count := service.UserList(r)

	response.Success(ctx, "", api.UserListResponse{
		Total: count,
		Data:  users,
	})
}

func GetUser(ctx *gin.Context) {
	var r api.GetUserRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	user := service.GetUser(r.Id)
	response.Success(ctx, "", user)
}

func EditUser(ctx *gin.Context) {
	var r api.EditUserRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	err = service.EditUser(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "修改成功")
}
