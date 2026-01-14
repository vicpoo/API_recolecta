package application

import "github.com/vicpoo/API_recolecta/src/RutaCamion/domain/ports"

type ExistsRutaCamionByIDUseCase struct {
	repo ports.RutaCamionRepository
}

func NewExistsRutaCamionByIDUseCase(repo ports.RutaCamionRepository) *ExistsRutaCamionByIDUseCase {
	return &ExistsRutaCamionByIDUseCase{repo: repo}
}

func (uc *ExistsRutaCamionByIDUseCase) Execute(id int32) (bool, error) {
	return uc.repo.ExistsByID(id)
}
