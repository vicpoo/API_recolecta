package core

import (
	"github.com/gin-gonic/gin"
)

func RegisterHealthEndpoint(engine *gin.Engine) {
	engine.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"message": "Backend is running",
		})
	})
}

