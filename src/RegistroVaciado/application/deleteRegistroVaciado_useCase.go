package application

import "github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/ports"

type DeleteRegistroVaciadoUseCase struct {
	repo ports.RegistroVaciadoRepository
}

func NewDeleteRegistroVaciadoUseCase(repo ports.RegistroVaciadoRepository) *DeleteRegistroVaciadoUseCase {
	return &DeleteRegistroVaciadoUseCase{repo: repo}
}

func (uc *DeleteRegistroVaciadoUseCase) Execute(id int32) error {
	return uc.repo.Delete(id)
}
