package application

import "github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"

type ListMisAlertas struct {
	repo domain.AlertaUsuarioRepository
}

func NewListMisAlertas(repo domain.AlertaUsuarioRepository) *ListMisAlertas {
	return &ListMisAlertas{repo: repo}
}

func (uc *ListMisAlertas) Execute(usuarioID int) ([]domain.AlertaUsuario, error) {
	return uc.repo.GetByUsuario(usuarioID)
}
