package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isnotvinicius/gopportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	// Get the id from the query parameters and validates it
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find the opening and send an error when not found
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
