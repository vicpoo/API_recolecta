package config

import "os"

type Config struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
	// Aqui se pueden anadir validaciones (ej. que no esten vacios)
	return cfg, nil
}
