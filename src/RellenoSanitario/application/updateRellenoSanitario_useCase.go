package application

import (
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/entities"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/ports"
)

type UpdateRellenoSanitarioUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewUpdateRellenoSanitarioUseCase(repo ports.RellenoSanitarioRepository) *UpdateRellenoSanitarioUseCase {
	return &UpdateRellenoSanitarioUseCase{repo}
}

func (uc *UpdateRellenoSanitarioUseCase) Execute(id int32, r *entities.RellenoSanitario) (*entities.RellenoSanitario, error) {
	return uc.repo.Update(id,r)
}
