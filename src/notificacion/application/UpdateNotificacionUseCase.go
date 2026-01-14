//UpdateNotificacionUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type UpdateNotificacionUseCase struct {
	repo repositories.INotificacion
}

func NewUpdateNotificacionUseCase(repo repositories.INotificacion) *UpdateNotificacionUseCase {
	return &UpdateNotificacionUseCase{repo: repo}
}

func (uc *UpdateNotificacionUseCase) Run(notificacion *entities.Notificacion) (*entities.Notificacion, error) {
	// Validar según el tipo antes de actualizar
	if err := notificacion.ValidarSegunTipo(); err != nil {
		return nil, err
	}
	
	err := uc.repo.Update(notificacion)
	if err != nil {
		return nil, err
	}
	return notificacion, nil
}