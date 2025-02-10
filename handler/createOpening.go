package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/isnotvinicius/gopportunities/schemas"
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

	// Maps the request values to a variable to avoid any data we don't want
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error while creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
