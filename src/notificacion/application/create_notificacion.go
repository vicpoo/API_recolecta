package application

import (
	"time"

	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type CreateNotificacion struct {
	repo domain.NotificacionRepository
}

func NewCreateNotificacion(repo domain.NotificacionRepository) *CreateNotificacion {
	return &CreateNotificacion{repo}
}

func (uc *CreateNotificacion) Execute(n *domain.Notificacion) error {
	n.Activa = true
	n.CreatedAt = time.Now()
	return uc.repo.Create(n)
}
