package controller

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/app/service"
	"data-export/pkg/g"
	"data-export/pkg/jwt"
	"data-export/pkg/response"
	"data-export/pkg/validator"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var r api.LoginRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	user, err := service.Login(r)
	if err != nil {
		response.Error(ctx, "用户名或密码错误")
		return
	}

	token, err := jwt.Login(int(user.Id))
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, "登录成功", api.LoginResponse{
		Token: token,
	})
}

func RefreshToken(ctx *gin.Context) {
	var r api.RefreshTokenRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		response.Error(ctx, "", validator.ProcessErr(r, err))
		return
	}

	token, err := jwt.RefreshToken(r.Token)
	if err != nil {
		response.Json(ctx, 100, err.Error())
		return
	}
	response.Success(ctx, "刷新成功", api.RefreshTokenResponse{
		Token: token,
	})
}

func Logout(ctx *gin.Context) {
	value, exists := ctx.Get("claims")
	if claims, ok := value.(*jwt.Claims); ok && exists {
		g.Redis().HDel(ctx, jwt.Config.RedisKey, claims.ID)
	}
	response.Success(ctx, "退出成功")
}

func GetUserInfo(ctx *gin.Context) {
	value, _ := ctx.Get("jwtUser")
	user := value.(model.User)
	response.Success(ctx, "", api.GetUserInfoResponse{
		Id:       user.Id,
		Username: user.Username,
		Menu:     service.UsesMenuList(ctx, nil, 0),
	})
}
