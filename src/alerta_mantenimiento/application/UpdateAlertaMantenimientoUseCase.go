//UpdateAlertaMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type UpdateAlertaMantenimientoUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewUpdateAlertaMantenimientoUseCase(repo repositories.IAlertaMantenimiento) *UpdateAlertaMantenimientoUseCase {
	return &UpdateAlertaMantenimientoUseCase{repo: repo}
}

func (uc *UpdateAlertaMantenimientoUseCase) Run(alertaMantenimiento *entities.AlertaMantenimiento) (*entities.AlertaMantenimiento, error) {
	err := uc.repo.Update(alertaMantenimiento)
	if err != nil {
		return nil, err
	}
	return alertaMantenimiento, nil
}