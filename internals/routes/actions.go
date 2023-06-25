package routes

import (
	"github.com/bllyanos/saved-actions/internals/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupActionsRoute(g *gin.Engine, db *gorm.DB) {
	group := g.Group("/actions")
	group.GET("/", handlers.GetAllActions(db))
	group.POST("/http/", handlers.CreateHttpAction(db))
}
