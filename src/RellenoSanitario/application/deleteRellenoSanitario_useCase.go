package application

import "github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/ports"

type DeleteRellenoSanitarioUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewDeleteRellenoSanitarioUseCase(repo ports.RellenoSanitarioRepository) *DeleteRellenoSanitarioUseCase {
	return &DeleteRellenoSanitarioUseCase{repo}
}

func (uc *DeleteRellenoSanitarioUseCase) Execute(id int) error {
	return uc.repo.Delete(id)
}
