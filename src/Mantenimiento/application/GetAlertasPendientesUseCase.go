//GetAlertasPendientesUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAlertasPendientesUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAlertasPendientesUseCase(repo repositories.IAlertaMantenimiento) *GetAlertasPendientesUseCase {
	return &GetAlertasPendientesUseCase{repo: repo}
}

func (uc *GetAlertasPendientesUseCase) Run() ([]entities.AlertaMantenimiento, error) {
	return uc.repo.GetPendientes()
}
