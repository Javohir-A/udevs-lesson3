package main

import (
	"github.com/udevs/lesson3/config"
	"github.com/udevs/lesson3/pkg/logger"
)

func main() {
	logger.Initialize()
	log := logger.GetLogger()

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to ")
	}
}
