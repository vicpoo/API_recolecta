//CreateRegistroMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type CreateRegistroMantenimientoUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewCreateRegistroMantenimientoUseCase(repo repositories.IRegistroMantenimiento) *CreateRegistroMantenimientoUseCase {
	return &CreateRegistroMantenimientoUseCase{repo: repo}
}

func (uc *CreateRegistroMantenimientoUseCase) Run(registro *entities.RegistroMantenimiento) (*entities.RegistroMantenimiento, error) {
	err := uc.repo.Save(registro)
	if err != nil {
		return nil, err
	}
	return registro, nil
}