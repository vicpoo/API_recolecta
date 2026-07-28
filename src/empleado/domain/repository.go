package domain

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/empleado/domain/entities"
)

type EmpleadoRepository interface {
	Create(ctx context.Context, empleado *entities.Empleado) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Empleado, error)
	List(ctx context.Context) ([]entities.Empleado, error)
	Update(ctx context.Context, empleado *entities.Empleado) error
	Delete(ctx context.Context, id int) error
	FindByMail(ctx context.Context, mail string) (*entities.Empleado, error)
	FindByUsername(ctx context.Context, username string) (*entities.Empleado, error)
	FindByMailOrUsername(ctx context.Context, value string) (*entities.Empleado, error)
}
