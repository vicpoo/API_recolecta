package domain

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/usuario/domain/entities"
)

type UsuarioRepository interface {
	Create(ctx context.Context, u *entities.Usuario) (int, error)
	GetByID(ctx context.Context, id int) (*entities.Usuario, error)
	List(ctx context.Context) ([]entities.Usuario, error)
	Delete(ctx context.Context, id int) error

	FindByEmail(ctx context.Context, email string) (*entities.Usuario, error)
}
