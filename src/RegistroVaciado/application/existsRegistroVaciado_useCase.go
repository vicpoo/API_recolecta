package application

import "github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/ports"

type ExistsRegistroVaciadoUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewExistsRegistroVaciadoUseCase(
	repo ports.RegistroVaciadoRepository,
) *ExistsRegistroVaciadoUseCase {
	return &ExistsRegistroVaciadoUseCase{repo: repo}
}

func (uc *ExistsRegistroVaciadoUseCase) Execute(id int32) (bool, error) {
	return uc.repo.ExistsByID(id)
}
