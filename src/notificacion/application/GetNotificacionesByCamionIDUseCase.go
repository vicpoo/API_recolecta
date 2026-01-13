//GetNotificacionesByCamionIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByCamionIDUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByCamionIDUseCase(repo repositories.INotificacion) *GetNotificacionesByCamionIDUseCase {
	return &GetNotificacionesByCamionIDUseCase{repo: repo}
}

func (uc *GetNotificacionesByCamionIDUseCase) Run(camionID int32) ([]entities.Notificacion, error) {
	return uc.repo.GetByCamionID(camionID)
}