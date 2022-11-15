package middlewares

import (
	"data-export/app/model"
	"data-export/pkg/g"
	"data-export/pkg/jwt"
	"data-export/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
	j "github.com/golang-jwt/jwt/v4"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, _ := getToken(ctx)
		claims, err := jwt.Parse(token)
		if err != nil {
			if errors.Is(err, j.ErrTokenExpired) {
				response.Json(ctx, 101, err.Error())
			} else {
				response.Json(ctx, 100, err.Error())
			}
			ctx.Abort()
			return
		}

		var user model.User
		g.DB().Where("id = ?", claims.ID).First(&user)
		ctx.Set("jwtUser", user)
		ctx.Set("claims", claims)

		ctx.Next()
	}
}

func getToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", j.ErrInvalidKey
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", j.ErrInvalidKey
	}
	return parts[1], nil
}
