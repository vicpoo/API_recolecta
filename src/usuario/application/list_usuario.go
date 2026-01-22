package application

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"github.com/vicpoo/API_recolecta/src/usuario/domain/entities"
)

type ViewAllUser struct {
	repo domain.UsuarioRepository
}

func NewViewAllUser(repo domain.UsuarioRepository) *ViewAllUser {
	return &ViewAllUser{repo: repo}
}

func (uc *ViewAllUser) Execute(ctx context.Context) ([]entities.Usuario, error) {
	return uc.repo.List(ctx)
}
