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
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
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

		config.MaxConns = 25
		config.MinConns = 5
		config.MaxConnLifetime = time.Hour
		config.MaxConnIdleTime = 30 * time.Minute
		config.HealthCheckPeriod = time.Minute

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pool, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			err = fmt.Errorf("error creating pool: %w", err)
			return
		}

		if pingErr := pool.Ping(ctx); pingErr != nil {
			pool.Close()
			err = fmt.Errorf("error pinging database: %w", pingErr)
			return
		}
	})

	return pool, err
}

func ClosePool() {
    if pool != nil {
        pool.Close()
    }
}


func GetBD() *pgxpool.Pool {
    db, err := ConnectPostgres()
    if err != nil {
        panic(fmt.Sprintf("Error al conectar a la base de datos: %v", err))
    }
    return db
}
