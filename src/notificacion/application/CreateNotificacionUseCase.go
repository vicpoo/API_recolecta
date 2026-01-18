//CreateNotificacionUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type CreateNotificacionUseCase struct {
	repo repositories.INotificacion
}

func NewCreateNotificacionUseCase(repo repositories.INotificacion) *CreateNotificacionUseCase {
	return &CreateNotificacionUseCase{repo: repo}
}

func (uc *CreateNotificacionUseCase) Run(notificacion *entities.Notificacion) (*entities.Notificacion, error) {
	// Validar según el tipo antes de guardar
	if err := notificacion.ValidarSegunTipo(); err != nil {
		return nil, err
	}
	
	err := uc.repo.Save(notificacion)
	if err != nil {
		return nil, err
	}
	return notificacion, nil
}