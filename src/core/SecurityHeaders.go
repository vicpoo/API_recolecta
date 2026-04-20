package core

import "github.com/gin-gonic/gin"

// SecurityHeadersMiddleware agrega headers HTTP de seguridad a todas las respuestas.
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Evita que el navegador adivine el Content-Type
		c.Header("X-Content-Type-Options", "nosniff")

		// Evita que la app sea embebida en iframes (clickjacking)
		c.Header("X-Frame-Options", "DENY")

		// Protección básica contra XSS en navegadores antiguos
		c.Header("X-XSS-Protection", "1; mode=block")

		// Fuerza HTTPS durante 1 año (solo activa en producción con TLS)
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Restringe de dónde se pueden cargar recursos
		c.Header("Content-Security-Policy", "default-src 'self'")

		// Evita filtrar la URL de referencia a sitios externos
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// Deshabilita funciones del navegador que no necesitas
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		c.Next()
	}
}