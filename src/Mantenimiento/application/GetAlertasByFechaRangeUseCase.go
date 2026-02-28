//GetAlertasByFechaRangeUseCase.go
package application

import (
	"time"
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAlertasByFechaRangeUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAlertasByFechaRangeUseCase(repo repositories.IAlertaMantenimiento) *GetAlertasByFechaRangeUseCase {
	return &GetAlertasByFechaRangeUseCase{repo: repo}
}

func (uc *GetAlertasByFechaRangeUseCase) Run(fechaInicio, fechaFin time.Time) ([]entities.AlertaMantenimiento, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}