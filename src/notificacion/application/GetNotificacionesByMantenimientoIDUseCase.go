//GetNotificacionesByMantenimientoIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByMantenimientoIDUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByMantenimientoIDUseCase(repo repositories.INotificacion) *GetNotificacionesByMantenimientoIDUseCase {
	return &GetNotificacionesByMantenimientoIDUseCase{repo: repo}
}

func (uc *GetNotificacionesByMantenimientoIDUseCase) Run(mantenimientoID int32) ([]entities.Notificacion, error) {
	return uc.repo.GetByMantenimientoID(mantenimientoID)
}