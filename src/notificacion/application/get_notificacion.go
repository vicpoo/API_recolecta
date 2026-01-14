package application

import "github.com/vicpoo/API_recolecta/src/notificacion/domain"

type GetNotificacion struct {
	repo domain.NotificacionRepository
}

func NewGetNotificacion(repo domain.NotificacionRepository) *GetNotificacion {
	return &GetNotificacion{repo}
}

func (uc *GetNotificacion) Execute(id int) (*domain.Notificacion, error) {
	return uc.repo.GetByID(id)
}
