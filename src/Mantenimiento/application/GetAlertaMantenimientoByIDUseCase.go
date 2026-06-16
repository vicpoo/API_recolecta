//GetAlertaMantenimientoByIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAlertaMantenimientoByIDUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAlertaMantenimientoByIDUseCase(repo repositories.IAlertaMantenimiento) *GetAlertaMantenimientoByIDUseCase {
	return &GetAlertaMantenimientoByIDUseCase{repo: repo}
}

func (uc *GetAlertaMantenimientoByIDUseCase) Run(id int32) (*entities.AlertaMantenimiento, error) {
	return uc.repo.GetByID(id)
}