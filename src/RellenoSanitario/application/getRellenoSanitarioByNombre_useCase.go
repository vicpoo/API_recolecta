package application

import (
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/ports"
)

type GetRellenoSanitarioByNombreUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewGetRellenoSanitarioByNombreUseCase(
	repo ports.RellenoSanitarioRepository,
) *GetRellenoSanitarioByNombreUseCase {
	return &GetRellenoSanitarioByNombreUseCase{repo: repo}
}

func (uc *GetRellenoSanitarioByNombreUseCase) Execute(
	nombre string,
) ([]entities.RellenoSanitario, error) {
	return uc.repo.GetByNombre(nombre)
}
