// GetSeguimientosFallaCriticaByFallaIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type GetSeguimientosFallaCriticaByFallaIDUseCase struct {
	repo repositories.ISeguimientoFallaCritica
}

func NewGetSeguimientosFallaCriticaByFallaIDUseCase(repo repositories.ISeguimientoFallaCritica) *GetSeguimientosFallaCriticaByFallaIDUseCase {
	return &GetSeguimientosFallaCriticaByFallaIDUseCase{repo: repo}
}

func (uc *GetSeguimientosFallaCriticaByFallaIDUseCase) Run(fallaID int32) ([]entities.SeguimientoFallaCritica, error) {
	return uc.repo.GetByFallaID(fallaID)
}