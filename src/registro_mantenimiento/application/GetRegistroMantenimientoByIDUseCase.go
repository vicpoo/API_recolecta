//GetRegistroMantenimientoByIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type GetRegistroMantenimientoByIDUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewGetRegistroMantenimientoByIDUseCase(repo repositories.IRegistroMantenimiento) *GetRegistroMantenimientoByIDUseCase {
	return &GetRegistroMantenimientoByIDUseCase{repo: repo}
}

func (uc *GetRegistroMantenimientoByIDUseCase) Run(id int32) (*entities.RegistroMantenimiento, error) {
	return uc.repo.GetByID(id)
}