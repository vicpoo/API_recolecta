package main

import (
	"log"

	"github.com/vicpoo/API_recolecta/src/core"
)

func main() {
	if _, err := core.ConnectRedis(); err != nil {
		log.Fatalf("failed to initialize redis: %v", err)
	}
	defer func() {
		if closeErr := core.CloseRedis(); closeErr != nil {
			log.Printf("failed to close redis: %v", closeErr)
		}
	}()

	InitDependencies()
}

