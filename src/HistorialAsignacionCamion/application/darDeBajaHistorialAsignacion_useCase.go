package application

import "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"

type DarDeBajaHistorialAsignacionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewDarDeBajaHistorialAsignacionUseCase(repo ports.IHistorialAsignacionCamion) *DarDeBajaHistorialAsignacionUseCase {
	return &DarDeBajaHistorialAsignacionUseCase{repo: repo}
}

func (uc *DarDeBajaHistorialAsignacionUseCase) Run(idHistorial int32) error {
	return uc.repo.DarDeBaja(idHistorial)
}
