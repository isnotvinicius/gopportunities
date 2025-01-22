package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Runs the server
	router.Run(":8080")
}
