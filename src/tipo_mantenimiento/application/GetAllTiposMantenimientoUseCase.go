// GetAllTiposMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain/entities"
)

type GetAllTiposMantenimientoUseCase struct {
	repo repositories.ITipoMantenimiento
}

func NewGetAllTiposMantenimientoUseCase(repo repositories.ITipoMantenimiento) *GetAllTiposMantenimientoUseCase {
	return &GetAllTiposMantenimientoUseCase{repo: repo}
}

func (uc *GetAllTiposMantenimientoUseCase) Run() ([]entities.TipoMantenimiento, error) {
	todos, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	
	// Filtrar solo los no eliminados
	var activos []entities.TipoMantenimiento
	for _, tm := range todos {
		if !tm.GetEliminado() {
			activos = append(activos, tm)
		}
	}
	
	return activos, nil
}