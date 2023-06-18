package routes

import (
	"time"

	"github.com/bllyanos/saved-actions/internals/core"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupIndexRoute(g *gin.Engine, db *gorm.DB) {
	group := g.Group("/")
	group.GET("/", func(ctx *gin.Context) {

		appName := "saved-actions"
		currentServerTime := time.Now().Format(time.RFC3339)

		core.SendResponse(
			ctx,
			200,
			"ok",
			gin.H{
				"appName":    appName,
				"serverTime": currentServerTime,
			},
		)
	})
}
