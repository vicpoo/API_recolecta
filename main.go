package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/vicpoo/API_recolecta/src/bootstrap"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) > 1 {
		switch strings.ToLower(strings.TrimSpace(os.Args[1])) {
		case "seed", "bootstrap":
			if err := bootstrap.SeedAdmin(context.Background()); err != nil {
				fmt.Fprintf(os.Stderr, "bootstrap failed: %v\n", err)
				os.Exit(1)
			}
			return
		}
	}

	InitDependencies()
}
