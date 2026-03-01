package application

import (
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)

type ListRellenoSanitarioUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewListRellenoSanitarioUseCase(repo ports.RellenoSanitarioRepository) *ListRellenoSanitarioUseCase {
	return &ListRellenoSanitarioUseCase{repo}
}

func (uc *ListRellenoSanitarioUseCase) Execute() ([]entities.RellenoSanitario, error) {
	return uc.repo.ListAll()
}
