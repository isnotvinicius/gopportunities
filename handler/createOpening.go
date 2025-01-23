package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CREATE",
	})
}
