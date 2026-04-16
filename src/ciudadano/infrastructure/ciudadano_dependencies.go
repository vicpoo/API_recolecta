package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type CiudadanoDependencies struct{}

func NewCiudadanoDependencies(_ *pgxpool.Pool, _ *redis.Client) *CiudadanoDependencies {
	return &CiudadanoDependencies{}
}

func RegisterCiudadanoRoutes(_ *gin.Engine, _ *CiudadanoDependencies) {}
