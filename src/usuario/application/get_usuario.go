package application

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"github.com/vicpoo/API_recolecta/src/usuario/domain/entities"
)

type ViewOneUser struct {
	repo domain.UsuarioRepository
}

func NewViewOneUser(repo domain.UsuarioRepository) *ViewOneUser {
	return &ViewOneUser{repo: repo}
}

func (uc *ViewOneUser) Execute(ctx context.Context, id int) (*entities.Usuario, error) {
	return uc.repo.GetByID(ctx, id)
}
