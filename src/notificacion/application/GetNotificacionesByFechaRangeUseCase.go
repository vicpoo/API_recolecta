//GetNotificacionesByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByFechaRangeUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByFechaRangeUseCase(repo repositories.INotificacion) *GetNotificacionesByFechaRangeUseCase {
	return &GetNotificacionesByFechaRangeUseCase{repo: repo}
}

func (uc *GetNotificacionesByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.Notificacion, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}