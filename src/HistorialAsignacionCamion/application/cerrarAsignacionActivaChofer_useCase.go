package application

import "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"

type CerrarAsignacionActivaChoferUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewCerrarAsignacionActivaChoferUseCase(repo ports.IHistorialAsignacionCamion) *CerrarAsignacionActivaChoferUseCase {
	return &CerrarAsignacionActivaChoferUseCase{repo: repo}
}

func (uc *CerrarAsignacionActivaChoferUseCase) Run(choferId int32) error {
	return uc.repo.CerrarAsignacionActivaChofer(choferId)
}
