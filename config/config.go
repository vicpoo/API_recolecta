package config

import "os"

type Config struct {
	RedisHost                string
	RedisPort                string
	RedisPassword            string
	FCMCredentialsFile       string
	ModeloReportesURL        string
	ClasificadorURL          string
	AnomaliaCreadaWebhookURL string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		FCMCredentialsFile: os.Getenv("FCM_CREDENTIALS_FILE"),
		// Nombres de servicio de docker.compose.yml: alcanzables por DNS interno
		// de Docker dentro de app_internal_net, tanto en dev como en despliegue
		// real (mismo compose stack). Si no estan seteadas, caen a esos defaults.
		ModeloReportesURL: getEnvOrDefault("MODELO_REPORTES_URL", "http://modelo_reportes:8000"),
		ClasificadorURL:   getEnvOrDefault("CLASIFICADOR_URL", "http://clasificador_reportes:8001"),
		// Webhook externo (equipo del algoritmo genetico de rutas) que se
		// notifica cuando se crea una anomalia con coordenadas. Todavia no
		// esta desplegado en ningun ambiente real -- por eso el default cae a
		// localhost:8004. Cuando lo desplieguen, solo hay que setear esta
		// variable de entorno (aqui y en docker-compose/.env de prod), no
		// tocar codigo.
		AnomaliaCreadaWebhookURL: getEnvOrDefault("ANOMALIA_CREADA_WEBHOOK_URL", "http://localhost:8004/anomalia_creada"),
	}
	// Aqui se pueden anadir validaciones (ej. que no esten vacios)
	return cfg, nil
}

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
