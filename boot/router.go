package boot

import (
	"data-export/pkg/config"
	"data-export/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type route struct {
	Port int
}

var Route = route{}

func (*route) Init() {
	config.UnmarshalKey("router", &Route)

	err := SetRouter().Run(fmt.Sprintf(":%v", Route.Port))
	if err != nil {
		panic(err)
	}
}

// SetRouter 设置路由与服务启动分开方便单元测试
func SetRouter() *gin.Engine {
	r := gin.Default()
	routers.Init(r)
	setup404Handler(r)
	return r
}

// 处理404请求
func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "not found",
		})
	})
}
