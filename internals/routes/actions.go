package routes

import (
	"fmt"
	"log"
	"time"

	"github.com/bllyanos/saved-actions/internals/core"
	"github.com/bllyanos/saved-actions/internals/dtos"
	"github.com/bllyanos/saved-actions/internals/models"
	"github.com/bllyanos/saved-actions/internals/models/subtypes/acttype"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupActionsRoute(g *gin.Engine, db *gorm.DB) {
	group := g.Group("/actions")

	group.GET("/", handleGet(db))
	group.POST("/http/", handleCreateHttpAction(db))
}

func handleGet(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		actions := []models.Action{
			{
				Name:       "Build Frontend",
				ActionType: acttype.Http,
				CreatedAt:  time.Now(),
			},
			{
				Name:       "Build Backend",
				ActionType: acttype.Docker,
				CreatedAt:  time.Now(),
			},
		}
		core.SendResponse(ctx, 200, "ok", actions)
	}
}

func handleCreateHttpAction(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dtos.CreateHttpAction

		if err := ctx.ShouldBindJSON(&dto); err != nil {
			log.Println(err.Error())
			core.SendResponse(ctx, 400, "bad request", nil)
			return
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			newAction := models.Action{
				Name:       dto.Name,
				ActionType: acttype.Http,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}

			if res := tx.Create(&newAction); res.Error != nil {
				return res.Error
			}

			newHttpActionDetail := models.HttpActionDetail{
				Url:       dto.Url,
				Body:      dto.Body,
				Method:    dto.Method,
				Headers:   dto.Headers,
				TimeoutMs: dto.TimeoutMs,
				ActionID:  newAction.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if res := tx.Create(&newHttpActionDetail); res.Error != nil {
				return res.Error
			}

			newActionParameters := []models.ActionParameter{}

			for _, param := range dto.Parameters {
				newActionParameters = append(newActionParameters, models.ActionParameter{
					ActionID:  newAction.ID,
					Name:      param.Name,
					Required:  param.Required,
					Default:   param.Default,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
			}

			tx.Create(&newActionParameters)

			return nil
		})

		if err != nil {
			fmt.Println("create error:", err.Error())
			core.SendResponse(ctx, 500, "internal server error", nil)
			return
		}

		core.SendResponse(ctx, 200, "ok", nil)
	}
}
