package ports

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/entities"
)

type CiudadanoPostgresRepository interface {
	Create(ctx context.Context, u *entities.CiudadanoPostgres) (int, error)
	FindByEmail(ctx context.Context, email string) (*entities.CiudadanoPostgres, error)
	FindByID(ctx context.Context, id int) (*entities.CiudadanoPostgres, error)
}
