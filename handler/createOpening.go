package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOpeningHandler(ctx *gin.Context) {
	// Initialize request with data from the body mapped by the struct
	request := CreateOpeningRequest{}

	// bind the JSON body to the request variable using gin context
	ctx.BindJSON(&request)

	// Validate the request
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		return
	}
}
