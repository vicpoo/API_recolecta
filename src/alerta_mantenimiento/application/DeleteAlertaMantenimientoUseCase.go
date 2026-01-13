// DeleteAlertaMantenimientoUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"

type DeleteAlertaMantenimientoUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewDeleteAlertaMantenimientoUseCase(repo repositories.IAlertaMantenimiento) *DeleteAlertaMantenimientoUseCase {
	return &DeleteAlertaMantenimientoUseCase{repo: repo}
}

func (uc *DeleteAlertaMantenimientoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}