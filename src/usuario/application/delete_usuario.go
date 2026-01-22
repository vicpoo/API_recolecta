package application

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
)

type DeleteUser struct {
	repo domain.UsuarioRepository
}

func NewDeleteUser(repo domain.UsuarioRepository) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (uc *DeleteUser) Execute(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
