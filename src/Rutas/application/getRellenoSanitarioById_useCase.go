package application

import (
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)

type GetRellenoSanitarioByIdUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewGetRellenoSanitarioByIdUseCase(repo ports.RellenoSanitarioRepository) *GetRellenoSanitarioByIdUseCase {
	return &GetRellenoSanitarioByIdUseCase{repo}
}

func (uc *GetRellenoSanitarioByIdUseCase) Execute(id int32) (*entities.RellenoSanitario, error) {
	return uc.repo.GetByID(id)
}
