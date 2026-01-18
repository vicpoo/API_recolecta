// GetTipoMantenimientoByIDUseCase.go
package application

import (
	"errors"
	repositories "github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain/entities"
)

type GetTipoMantenimientoByIDUseCase struct {
	repo repositories.ITipoMantenimiento
}

func NewGetTipoMantenimientoByIDUseCase(repo repositories.ITipoMantenimiento) *GetTipoMantenimientoByIDUseCase {
	return &GetTipoMantenimientoByIDUseCase{repo: repo}
}

func (uc *GetTipoMantenimientoByIDUseCase) Run(id int32) (*entities.TipoMantenimiento, error) {
	tipoMantenimiento, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	// Verificar si está eliminado
	if tipoMantenimiento.GetEliminado() {
		return nil, errors.New("el tipo de mantenimiento ha sido eliminado")
	}
	
	return tipoMantenimiento, nil
}