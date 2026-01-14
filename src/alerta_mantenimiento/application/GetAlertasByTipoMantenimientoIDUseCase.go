//GetAlertasByTipoMantenimientoIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
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

