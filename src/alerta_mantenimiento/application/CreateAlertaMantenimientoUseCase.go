//CreateAlertaMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type CreateAlertaMantenimientoUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewCreateAlertaMantenimientoUseCase(repo repositories.IAlertaMantenimiento) *CreateAlertaMantenimientoUseCase {
	return &CreateAlertaMantenimientoUseCase{repo: repo}
}

func (uc *CreateAlertaMantenimientoUseCase) Run(alertaMantenimiento *entities.AlertaMantenimiento) (*entities.AlertaMantenimiento, error) {
	err := uc.repo.Save(alertaMantenimiento)
	if err != nil {
		return nil, err
	}
	return alertaMantenimiento, nil
}