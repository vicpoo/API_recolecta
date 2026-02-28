//GetAlertasByCamionIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAlertasByCamionIDUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAlertasByCamionIDUseCase(repo repositories.IAlertaMantenimiento) *GetAlertasByCamionIDUseCase {
	return &GetAlertasByCamionIDUseCase{repo: repo}
}

func (uc *GetAlertasByCamionIDUseCase) Run(camionID int32) ([]entities.AlertaMantenimiento, error) {
	return uc.repo.GetByCamionID(camionID)
}
