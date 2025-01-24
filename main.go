package main

import (
	"github.com/isnotvinicius/gopportunities/config"
	"github.com/isnotvinicius/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main package")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config init resulted in error: %v", err)
		return
	}

	// Initialize the router package
	router.Initialize()
}
