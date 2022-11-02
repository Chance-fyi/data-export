package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(ctx *gin.Context, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}
	ctx.JSON(http.StatusOK, jsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

func Success(ctx *gin.Context, message string, data ...interface{}) {
	Json(ctx, 0, message, data...)
}

func Error(ctx *gin.Context, message string, data ...interface{}) {
	Json(ctx, 1, message, data...)
}
