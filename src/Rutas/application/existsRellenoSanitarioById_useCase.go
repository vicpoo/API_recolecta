package application

import "github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"

type ExistsRellenoSanitarioByIdUseCase struct {
	repo ports.RellenoSanitarioRepository
}

func NewExistsRellenoSanitarioByIdUseCase(
	repo ports.RellenoSanitarioRepository,
) *ExistsRellenoSanitarioByIdUseCase {
	return &ExistsRellenoSanitarioByIdUseCase{repo: repo}
}

func (uc *ExistsRellenoSanitarioByIdUseCase) Execute(
	id int32,
) (bool, error) {
	return uc.repo.ExistsByID(id)
}
