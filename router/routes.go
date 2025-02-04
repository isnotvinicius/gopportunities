package router

import (
	"github.com/gin-gonic/gin"

	"github.com/isnotvinicius/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.Init()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.PostOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
