//GetRegistrosByCoordinadorIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type GetRegistrosByCoordinadorIDUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewGetRegistrosByCoordinadorIDUseCase(repo repositories.IRegistroMantenimiento) *GetRegistrosByCoordinadorIDUseCase {
	return &GetRegistrosByCoordinadorIDUseCase{repo: repo}
}

func (uc *GetRegistrosByCoordinadorIDUseCase) Run(coordinadorID int32) ([]entities.RegistroMantenimiento, error) {
	return uc.repo.GetByCoordinadorID(coordinadorID)
}