//GetRegistrosByFechaRangeUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type GetRegistrosByFechaRangeUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewGetRegistrosByFechaRangeUseCase(repo repositories.IRegistroMantenimiento) *GetRegistrosByFechaRangeUseCase {
	return &GetRegistrosByFechaRangeUseCase{repo: repo}
}

func (uc *GetRegistrosByFechaRangeUseCase) Run(fechaInicio, fechaFin string) ([]entities.RegistroMantenimiento, error) {
	return uc.repo.GetByFechaRange(fechaInicio, fechaFin)
}