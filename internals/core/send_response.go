package core

import "github.com/gin-gonic/gin"

func SendResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, gin.H{
		"message": message,
		"status":  status,
		"data":    data,
	})
}
