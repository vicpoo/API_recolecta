// UpdateTipoMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type UpdateTipoMantenimientoUseCase struct {
	repo repositories.ITipoMantenimiento
}

func NewUpdateTipoMantenimientoUseCase(repo repositories.ITipoMantenimiento) *UpdateTipoMantenimientoUseCase {
	return &UpdateTipoMantenimientoUseCase{repo: repo}
}

func (uc *UpdateTipoMantenimientoUseCase) Run(tipoMantenimiento *entities.TipoMantenimiento) (*entities.TipoMantenimiento, error) {
	// Obtener el registro actual para preservar el estado "eliminado"
	actual, err := uc.repo.GetByID(tipoMantenimiento.GetID())
	if err != nil {
		return nil, err
	}
	
	// Preservar el estado de eliminado
	tipoMantenimiento.SetEliminado(actual.GetEliminado())
	
	// Actualizar solo nombre y categoría
	actual.SetNombre(tipoMantenimiento.GetNombre())
	actual.SetCategoria(tipoMantenimiento.GetCategoria())
	
	err = uc.repo.Update(actual)
	if err != nil {
		return nil, err
	}
	return actual, nil
}