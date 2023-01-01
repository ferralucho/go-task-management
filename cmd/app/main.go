package main

import (
	"log"

	"github.com/ferralucho/go-task-management/config"
	"github.com/ferralucho/go-task-management/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
