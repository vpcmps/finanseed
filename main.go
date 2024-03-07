package main

import (
	"github.com/vpcmps/finanseed/config"
	"github.com/vpcmps/finanseed/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	//Initialize the router
	router.Initialize()
}
