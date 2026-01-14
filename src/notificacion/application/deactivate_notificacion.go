package application

import "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type DeactivateNotificacion struct {
	repo domain.NotificacionRepository
}

func NewDeactivateNotificacion(repo domain.NotificacionRepository) *DeactivateNotificacion {
	return &DeactivateNotificacion{repo}
}

func (uc *DeactivateNotificacion) Execute(id int) error {
	return uc.repo.Deactivate(id)
}
