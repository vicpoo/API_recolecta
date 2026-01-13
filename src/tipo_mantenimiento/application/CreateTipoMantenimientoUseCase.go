// CreateTipoMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain/entities"
)

type CreateTipoMantenimientoUseCase struct {
	repo repositories.ITipoMantenimiento
}

func NewCreateTipoMantenimientoUseCase(repo repositories.ITipoMantenimiento) *CreateTipoMantenimientoUseCase {
	return &CreateTipoMantenimientoUseCase{repo: repo}
}

func (uc *CreateTipoMantenimientoUseCase) Run(tipoMantenimiento *entities.TipoMantenimiento) (*entities.TipoMantenimiento, error) {
	err := uc.repo.Save(tipoMantenimiento)
	if err != nil {
		return nil, err
	}
	return tipoMantenimiento, nil
}