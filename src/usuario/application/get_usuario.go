package application

import "github.com/vicpoo/API_recolecta/src/usuario/domain"

type GetUsuario struct {
	repo domain.UsuarioRepository
}

func NewGetUsuario(repo domain.UsuarioRepository) *GetUsuario {
	return &GetUsuario{repo}
}

func (uc *GetUsuario) Execute(id int) (*domain.Usuario, error) {
	return uc.repo.GetByID(id)
}
