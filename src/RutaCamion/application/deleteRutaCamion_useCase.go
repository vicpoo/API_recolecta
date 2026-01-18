package application

import "github.com/vicpoo/API_recolecta/src/RutaCamion/domain/ports"

type DeleteRutaCamionUseCase struct {
	repo ports.RutaCamionRepository
}

func NewDeleteRutaCamionUseCase(repo ports.RutaCamionRepository) *DeleteRutaCamionUseCase {
	return &DeleteRutaCamionUseCase{repo: repo}
}

func (uc *DeleteRutaCamionUseCase) Execute(id int32) error {
	return uc.repo.Delete(id)
}
