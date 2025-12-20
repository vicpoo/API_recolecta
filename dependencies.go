package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vicpoo/API_recolecta/src/core"
)

//archivo para hacer las instancias de los controllers, casos de uso y repositories, etc.
func InitDependencies() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("error al cargar el .env")
	}

	engine := gin.Default()
	engine.Use(core.CORSMiddleware())
	
	engine.Run(":8080")
}