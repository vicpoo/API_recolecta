package application

import "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"

type DeleteHistorialAsignacionCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewDeleteHistorialAsignacionCamionUseCase(repo ports.IHistorialAsignacionCamion) *DeleteHistorialAsignacionCamionUseCase {
	return &DeleteHistorialAsignacionCamionUseCase{
		repo: repo, 
	}
}

func (uc *DeleteHistorialAsignacionCamionUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}