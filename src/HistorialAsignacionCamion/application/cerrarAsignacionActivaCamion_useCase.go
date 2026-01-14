package application

import "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/ports"

type CerrarAsignacionActivaCamionUseCase struct {
	repo ports.IHistorialAsignacionCamion
}

func NewCerrarAsignacionActivaCamionUseCase(repo ports.IHistorialAsignacionCamion) *CerrarAsignacionActivaCamionUseCase {
	return &CerrarAsignacionActivaCamionUseCase{
		repo: repo,
	}
}

// Cierra cualquier asignación ACTIVA del camión
func (uc *CerrarAsignacionActivaCamionUseCase) Run(camionId int32) error {
	return uc.repo.CerrarAsignacionActivaCamion(camionId)
}
