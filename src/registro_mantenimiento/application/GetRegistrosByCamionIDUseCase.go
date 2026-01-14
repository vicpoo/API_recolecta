//GetRegistrosByCamionIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type GetRegistrosByCamionIDUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewGetRegistrosByCamionIDUseCase(repo repositories.IRegistroMantenimiento) *GetRegistrosByCamionIDUseCase {
	return &GetRegistrosByCamionIDUseCase{repo: repo}
}

func (uc *GetRegistrosByCamionIDUseCase) Run(camionID int32) ([]entities.RegistroMantenimiento, error) {
	return uc.repo.GetByCamionID(camionID)
}