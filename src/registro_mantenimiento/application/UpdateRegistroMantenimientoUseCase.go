//UpdateRegistroMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type UpdateRegistroMantenimientoUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewUpdateRegistroMantenimientoUseCase(repo repositories.IRegistroMantenimiento) *UpdateRegistroMantenimientoUseCase {
	return &UpdateRegistroMantenimientoUseCase{repo: repo}
}

func (uc *UpdateRegistroMantenimientoUseCase) Run(registro *entities.RegistroMantenimiento) (*entities.RegistroMantenimiento, error) {
	err := uc.repo.Update(registro)
	if err != nil {
		return nil, err
	}
	return registro, nil
}