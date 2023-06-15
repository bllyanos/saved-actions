package routes

import (
	"time"

	"github.com/bllyanos/saved-actions/internals/core"
	"github.com/bllyanos/saved-actions/internals/models"
	"github.com/gin-gonic/gin"
)

func SetupActionsRoute(g *gin.Engine) {
	group := g.Group("/actions")
	group.GET("/", func(ctx *gin.Context) {

		actions := []models.Action{
			{
				Name:      "Build Frontend",
				CreatedAt: time.Now(),
			},
			{
				Name:      "Build Backend",
				CreatedAt: time.Now(),
			},
		}

		core.SendResponse(ctx, 200, "ok", actions)
	})
}
