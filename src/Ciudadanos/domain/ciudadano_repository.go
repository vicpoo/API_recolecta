package domain

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
)

type CiudadanoRepository interface {
	Create(ctx context.Context, ciudadano *entities.Ciudadano) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Ciudadano, error)
	List(ctx context.Context) ([]entities.Ciudadano, error)
	Update(ctx context.Context, ciudadano *entities.Ciudadano) error
	Delete(ctx context.Context, id int) error
	FindByEmail(ctx context.Context, email string) (*entities.Ciudadano, error)
	FindByAlias(ctx context.Context, alias string) (*entities.Ciudadano, error)
	FindByEmailOrAlias(ctx context.Context, value string) (*entities.Ciudadano, error)
}