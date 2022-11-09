package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/app"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateExport(ctx *gin.Context) {
	var r api.CreateExportRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	err = service.CreateExport(r, app.JwtUser(ctx).Id)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}
