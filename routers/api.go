package routers

import (
	"data-export/app/controller"
	"data-export/app/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/login", controller.Login)
		api.POST("/refreshToken", controller.RefreshToken)

		api.Use(middlewares.AuthJWT())
		{
			api.GET("/logout", controller.Logout)
			api.GET("/getUserInfo", controller.GetUserInfo)

			user := api.Group("/user")
			{
				user.POST("create", controller.CreateUser)
			}

			menu := api.Group("/menu")
			{
				menu.POST("create", controller.CreateMenu)
				menu.GET("list", controller.MenuList)
				menu.GET("selectTree", controller.MenuSelectTree)
			}
		}
	}
}
