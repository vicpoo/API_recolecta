package core

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
    pool *pgxpool.Pool
    once sync.Once
    err  error
)

func ConnectPostgres() (*pgxpool.Pool, error) {
    once.Do(func() {
        dsn := fmt.Sprintf(
            "postgres://%s:%s@%s:%s/%s",
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_NAME"),
        )

        config, parseErr := pgxpool.ParseConfig(dsn)
        if parseErr != nil {
            err = fmt.Errorf("error parsing DSN: %w", parseErr)
            return
        }

        // Configuraciones del pool
        config.MaxConns = 25                          // Máximo de conexiones
        config.MinConns = 5                           // Mínimo de conexiones
        config.MaxConnLifetime = time.Hour            // Tiempo de vida
        config.MaxConnIdleTime = 30 * time.Minute     // Tiempo inactivo
        config.HealthCheckPeriod = time.Minute        // Health check

        // Crear el pool
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        pool, err = pgxpool.NewWithConfig(ctx, config)
        if err != nil {
            err = fmt.Errorf("error creating pool: %w", err)
            return
        }

        // Verificar conexión
        if pingErr := pool.Ping(ctx); pingErr != nil {
            pool.Close()
            err = fmt.Errorf("error pinging database: %w", pingErr)
            return
        }
    })

    return pool, err
}

// ClosePool cierra el pool de conexiones
func ClosePool() {
    if pool != nil {
        pool.Close()
    }
}