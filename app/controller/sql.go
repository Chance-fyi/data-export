package controller

import (
	"data-export/app/api"
	"data-export/app/service"
	"data-export/pkg/app"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateSql(ctx *gin.Context) {
	var r api.CreateSqlRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.CreateSql(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "创建成功")
}

func SqlList(ctx *gin.Context) {
	var r api.SqlListRequest
	_ = ctx.ShouldBind(&r)

	sqls, count := service.SqlList(r)

	response.Success(ctx, "", api.SqlListResponse{
		Total: count,
		Data:  sqls,
	})
}

func GetSql(ctx *gin.Context) {
	var r api.GetSqlRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	sql := service.GetSql(r.Id)
	response.Success(ctx, "", sql)
}

func EditSql(ctx *gin.Context) {
	var r api.EditSqlRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.EditSql(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "修改成功")
}

func GetUserSql(ctx *gin.Context) {
	var r api.GetUserSqlRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	userIds := service.GetUserSql(r.Id)
	data := api.GetUserSqlResponse{
		Id:      r.Id,
		UserIds: []string{},
	}
	if userIds != nil {
		data.UserIds = userIds
	}

	response.Success(ctx, "", data)
}

func SetUserSql(ctx *gin.Context) {
	var r api.SetUserSqlRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	err = service.SetUserSql(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "设置成功")
}

func MySqlList(ctx *gin.Context) {
	var r api.MySqlListRequest
	_ = ctx.ShouldBind(&r)

	user := app.JwtUser(ctx)
	list, count := service.MySqlList(r, user)

	response.Success(ctx, "", api.MySqlListResponse{
		Total: count,
		Data:  list,
	})
}

func GetUserSqlName(ctx *gin.Context) {
	var r api.GetUserSqlNameRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	data := service.GetUserSqlName(r.Id)
	response.Success(ctx, "", data)
}

func SetUserSqlName(ctx *gin.Context) {
	var r api.SetUserSqlNameRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	err = service.SetUserSqlName(r)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "修改成功")
}

func GetDownloadSql(ctx *gin.Context) {
	var r api.GetDownloadSqlRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}
	data := service.GetDownloadSql(r.Id)
	response.Success(ctx, "", data)
}
