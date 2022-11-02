package controller

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/app/service"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateUser(ctx *gin.Context) {
	var r api.CreateUserRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.CreateUser(model.User{
		Username:   r.Username,
		Password:   r.Password,
		CreateTime: time.Now(),
	})
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}
