package main

import (
	"github.com/AndersonMeloo/Go.git/config"
	"github.com/AndersonMeloo/Go.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
