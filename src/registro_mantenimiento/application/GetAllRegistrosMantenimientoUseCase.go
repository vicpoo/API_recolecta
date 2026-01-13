//GetAllRegistrosMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type GetAllRegistrosMantenimientoUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewGetAllRegistrosMantenimientoUseCase(repo repositories.IRegistroMantenimiento) *GetAllRegistrosMantenimientoUseCase {
	return &GetAllRegistrosMantenimientoUseCase{repo: repo}
}

func (uc *GetAllRegistrosMantenimientoUseCase) Run() ([]entities.RegistroMantenimiento, error) {
	return uc.repo.GetAll()
}