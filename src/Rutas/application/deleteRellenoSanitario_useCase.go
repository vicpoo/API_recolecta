package application

import "github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"

type DeleteRellenoSanitarioUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewDeleteRellenoSanitarioUseCase(repo ports.RellenoSanitarioRepository) *DeleteRellenoSanitarioUseCase {
	return &DeleteRellenoSanitarioUseCase{repo}
}

func (uc *DeleteRellenoSanitarioUseCase) Execute(id int32) error {
	return uc.repo.Delete(id)
}
