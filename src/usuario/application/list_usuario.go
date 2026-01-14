package application

import "github.com/vicpoo/API_recolecta/src/usuario/domain"

type ListUsuarios struct {
	repo domain.UsuarioRepository
}

func NewListUsuarios(repo domain.UsuarioRepository) *ListUsuarios {
	return &ListUsuarios{repo}
}

func (uc *ListUsuarios) Execute() ([]domain.Usuario, error) {
	return uc.repo.GetAll()
}
