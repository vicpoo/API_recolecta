//DeleteRegistroMantenimientoUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"

type DeleteRegistroMantenimientoUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewDeleteRegistroMantenimientoUseCase(repo repositories.IRegistroMantenimiento) *DeleteRegistroMantenimientoUseCase {
	return &DeleteRegistroMantenimientoUseCase{repo: repo}
}

func (uc *DeleteRegistroMantenimientoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}