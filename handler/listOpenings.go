package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isnotvinicius/gopportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error while listing openings on database")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
