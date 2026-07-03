package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) > 1 {
		switch strings.ToLower(strings.TrimSpace(os.Args[1])) {
		case "seed", "bootstrap":
			fmt.Fprintln(os.Stderr, "bootstrap de admin deshabilitado: gestiona usuarios desde scripts/init/.env")
			os.Exit(1)
		}
	}

	InitDependencies()
}
