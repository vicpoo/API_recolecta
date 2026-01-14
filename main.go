package main

import "log"

func main() {
	if err := InitDependencies(); err != nil {
		log.Fatal(err)
	}
}