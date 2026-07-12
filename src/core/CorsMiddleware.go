package core

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	originsEnv := os.Getenv("CORS_ALLOWED_ORIGINS")
	if originsEnv == "" {
		originsEnv = os.Getenv("CORS_ORIGINS")
	}

	var origins []string
	for _, origin := range strings.Split(originsEnv, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			origins = append(origins, origin)
		}
	}

	// Si no hay orígenes configurados, permitir todos dinámicamente para soportar credenciales en desarrollo
	if len(origins) == 0 {
		return cors.New(cors.Config{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
	}

	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

