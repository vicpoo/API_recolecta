package application

import "github.com/vicpoo/API_recolecta/src/usuario/domain"

type DeleteUsuario struct {
	repo domain.UsuarioRepository
}

func NewDeleteUsuario(repo domain.UsuarioRepository) *DeleteUsuario {
	return &DeleteUsuario{repo}
}

func (uc *DeleteUsuario) Execute(id int) error {
	return uc.repo.Delete(id)
}