//GetAlertasByTipoMantenimientoIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAlertasByTipoMantenimientoIDUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAlertasByTipoMantenimientoIDUseCase(repo repositories.IAlertaMantenimiento) *GetAlertasByTipoMantenimientoIDUseCase {
	return &GetAlertasByTipoMantenimientoIDUseCase{repo: repo}
}

func (uc *GetAlertasByTipoMantenimientoIDUseCase) Run(tipoID int32) ([]entities.AlertaMantenimiento, error) {
	return uc.repo.GetByTipoMantenimientoID(tipoID)
}

