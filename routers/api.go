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
				user.GET("list", controller.UserList)
				user.POST("create", controller.CreateUser)
				user.GET("get", controller.GetUser)
				user.POST("edit", controller.EditUser)
				user.GET("selectList", controller.UserSelectList)
			}

			menu := api.Group("/menu")
			{
				menu.POST("create", controller.CreateMenu)
				menu.POST("edit", controller.EditMenu)
				menu.GET("get", controller.GetMenu)
				menu.GET("list", controller.MenuList)
				menu.GET("selectTree", controller.MenuSelectTree)
			}

			role := api.Group("/role")
			{
				role.POST("create", controller.CreateRole)
				role.GET("list", controller.RoleList)
				role.GET("get", controller.GetRole)
				role.POST("edit", controller.EditRole)
				role.GET("selectList", controller.RoleSelectList)
			}

			database := api.Group("/database")
			{
				database.POST("create", controller.CreateDatabase)
				database.GET("list", controller.DatabaseList)
				database.GET("get", controller.GetDatabase)
				database.POST("edit", controller.EditDatabase)
			}

			sql := api.Group("/sql")
			{
				sql.POST("create", controller.CreateSql)
				sql.GET("list", controller.SqlList)
				sql.GET("get", controller.GetSql)
				sql.POST("edit", controller.EditSql)
				sql.GET("getUser", controller.GetUserSql)
				sql.POST("setUser", controller.SetUserSql)
			}
		}
	}
}
