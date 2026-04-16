package config

import "os"

type Config struct {
	RedisHost          string
	RedisPort          string
	RedisPassword      string
	FCMCredentialsFile string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		FCMCredentialsFile: os.Getenv("FCM_CREDENTIALS_FILE"),
	}
	// Aqui se pueden anadir validaciones (ej. que no esten vacios)
	return cfg, nil
}
