package actions

import (
	"github.com/bllyanos/saved-actions/internals/core"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleGetAllActions(db *gorm.DB) gin.HandlerFunc {
	handleGetAllActions := func(ctx *gin.Context) {
		core.SendResponse(ctx, 500, "not implemented", nil)
	}
	return handleGetAllActions
}
