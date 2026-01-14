//GetAllIncidenciasUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type GetAllIncidenciasUseCase struct {
	repo repositories.IIncidencia
}

func NewGetAllIncidenciasUseCase(repo repositories.IIncidencia) *GetAllIncidenciasUseCase {
	return &GetAllIncidenciasUseCase{repo: repo}
}

func (uc *GetAllIncidenciasUseCase) Run() ([]entities.Incidencia, error) {
	return uc.repo.GetAll()
}