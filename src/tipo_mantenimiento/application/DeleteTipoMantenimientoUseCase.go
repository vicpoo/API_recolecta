// DeleteTipoMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain"
)

type DeleteTipoMantenimientoUseCase struct {
	repo repositories.ITipoMantenimiento
}

func NewDeleteTipoMantenimientoUseCase(repo repositories.ITipoMantenimiento) *DeleteTipoMantenimientoUseCase {
	return &DeleteTipoMantenimientoUseCase{repo: repo}
}

func (uc *DeleteTipoMantenimientoUseCase) Run(id int32) error {
	// Obtener el registro primero
	tipoMantenimiento, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	// Marcar como eliminado (borrado lógico)
	tipoMantenimiento.MarcarEliminado()
	
	// Actualizar el registro
	return uc.repo.Update(tipoMantenimiento)
}