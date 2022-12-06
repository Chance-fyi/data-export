package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/app"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

func ExportList(ctx *gin.Context) {
	var r api.ExportListRequest
	_ = ctx.ShouldBind(&r)

	list, count := service.ExportList(r, app.JwtUser(ctx).Id)
	response.Success(ctx, "", api.ExportListResponse{
		Total: count,
		Data:  list,
	})
}

func ExportDownload(ctx *gin.Context) {
	var r api.ExportDownloadRequest
	_ = ctx.ShouldBind(&r)
	e := service.ExportDownload(r)
	_, err := os.Open(e.Path)
	if e.Path == "" || err != nil {
		ctx.Redirect(http.StatusFound, "/404")
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%v.xlsx", e.Filename))
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(e.Path)
	return
}
