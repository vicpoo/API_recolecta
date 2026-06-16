//GetAlertasAtendidasUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type GetAlertasAtendidasUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAlertasAtendidasUseCase(repo repositories.IAlertaMantenimiento) *GetAlertasAtendidasUseCase {
	return &GetAlertasAtendidasUseCase{repo: repo}
}

func (uc *GetAlertasAtendidasUseCase) Run() ([]entities.AlertaMantenimiento, error) {
	return uc.repo.GetAtendidas()
}