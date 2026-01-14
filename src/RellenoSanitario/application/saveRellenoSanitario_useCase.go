package application

import (
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/ports"
)

type SaveRellenoSanitarioUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewSaveRellenoSanitarioUseCase(repo ports.RellenoSanitarioRepository) *SaveRellenoSanitarioUseCase {
	return &SaveRellenoSanitarioUseCase{repo}
}

func (uc *SaveRellenoSanitarioUseCase) Execute(r *entities.RellenoSanitario) (*entities.RellenoSanitario, error) {
	return uc.repo.Save(r)
}
