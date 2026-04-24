package application

import "github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"

type MarcarLeida struct {
	repo domain.AlertaUsuarioRepository
}

func NewMarcarLeida(repo domain.AlertaUsuarioRepository) *MarcarLeida {
	return &MarcarLeida{repo: repo}
}

func (uc *MarcarLeida) Execute(alertaID int, usuarioID int) error {
	return uc.repo.MarkAsRead(alertaID, usuarioID)
}
