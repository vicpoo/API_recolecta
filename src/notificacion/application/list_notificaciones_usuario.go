package application

import "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type ListNotificacionesUsuario struct {
	repo domain.NotificacionRepository
}

func NewListNotificacionesUsuario(repo domain.NotificacionRepository) *ListNotificacionesUsuario {
	return &ListNotificacionesUsuario{repo}
}

func (uc *ListNotificacionesUsuario) Execute(usuarioID int) ([]domain.Notificacion, error) {
	return uc.repo.GetByUsuario(usuarioID)
}
