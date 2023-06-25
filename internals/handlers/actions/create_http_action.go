package actions

import (
	"log"

	"github.com/bllyanos/saved-actions/internals/core"
	"github.com/bllyanos/saved-actions/internals/dtos"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleCreateHttpAction(db *gorm.DB) gin.HandlerFunc {
	createHttpAction := func(ctx *gin.Context) {
		var dto dtos.CreateHttpAction

		if err := ctx.ShouldBindJSON(&dto); err != nil {
			log.Println("body binding error:", err.Error())
			core.SendResponse(ctx, 400, "bad request", nil)
			return
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			newAction := dto.GetAction()
			if res := tx.Create(&newAction); res.Error != nil {
				return res.Error
			}

			newHttpActionDetail := dto.GetHttpActionDetail(newAction.ID)
			if res := tx.Create(&newHttpActionDetail); res.Error != nil {
				return res.Error
			}

			newActionParameters := dto.GetActionParameter(newAction.ID)
			if res := tx.Create(&newActionParameters); res.Error != nil {
				return res.Error
			}

			return nil
		})

		if err != nil {
			log.Println("create http action error", err.Error())
			core.SendResponse(ctx, 500, "internal server error", nil)
			return
		}

		core.SendResponse(ctx, 200, "ok", nil)
	}
	return createHttpAction
}
