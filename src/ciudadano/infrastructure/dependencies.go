package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"
	goredis "github.com/redis/go-redis/v9"
	ciudadanoApp "github.com/vicpoo/API_recolecta/src/ciudadano/application"
	ciudadanoHttp "github.com/vicpoo/API_recolecta/src/ciudadano/infrastructure/http"
	ciudadanoPostgres "github.com/vicpoo/API_recolecta/src/ciudadano/infrastructure/postgres"
	ciudadanoRedis "github.com/vicpoo/API_recolecta/src/ciudadano/infrastructure/redis"
)

type CiudadanoDependencies struct {
	Controller *ciudadanoHttp.CiudadanoController
}

func NewCiudadanoDependencies(db *pgxpool.Pool, rdb *goredis.Client) *CiudadanoDependencies {
	postgresRepo := ciudadanoPostgres.NewPostgresCiudadanoRepository(db)
	redisRepo := ciudadanoRedis.NewRedisCiudadanoRepository(rdb)

	registerUC := ciudadanoApp.NewRegisterCiudadanoUseCase(postgresRepo, redisRepo)
	updateCoordUC := ciudadanoApp.NewUpdateCoordinatesUseCase(redisRepo, postgresRepo)

	controller := ciudadanoHttp.NewCiudadanoController(registerUC, updateCoordUC)

	return &CiudadanoDependencies{Controller: controller}
}
