package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateDatabase(ctx *gin.Context) {
	var r api.CreateDatabaseRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.CreateDatabase(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}

func DatabaseList(ctx *gin.Context) {
	var r api.DatabaseListRequest
	_ = ctx.ShouldBind(&r)

	databases, count := service.DatabaseList(r)

	response.Success(ctx, "", api.DatabaseListResponse{
		Total: count,
		Data:  databases,
	})
}

func GetDatabase(ctx *gin.Context) {
	var r api.GetDatabaseRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	database := service.GetDatabase(r.Id)
	response.Success(ctx, "", database)
}

func EditDatabase(ctx *gin.Context) {
	var r api.EditDatabaseRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.EditDatabase(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "修改成功")
}
