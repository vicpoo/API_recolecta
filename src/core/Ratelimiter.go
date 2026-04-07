package core

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	count    int
	resetAt  time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

// RateLimitMiddleware limita el número de peticiones por IP en memoria.
// maxRequests: número máximo de peticiones permitidas en la ventana de tiempo.
// window: duración de la ventana (ej: time.Minute).
func RateLimitMiddleware(maxRequests int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		entry, exists := clients[ip]

		if !exists || time.Now().After(entry.resetAt) {
			// Primera vez o ventana expirada — reiniciamos el contador
			clients[ip] = &client{
				count:   1,
				resetAt: time.Now().Add(window),
			}
			mu.Unlock()

			c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", maxRequests))
			c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", maxRequests-1))
			c.Next()
			return
		}

		entry.count++
		count := entry.count
		mu.Unlock()

		remaining := maxRequests - count
		if remaining < 0 {
			remaining = 0
		}

		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", maxRequests))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))

		if count > maxRequests {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "demasiadas solicitudes, intenta más tarde",
			})
			return
		}

		c.Next()
	}
}