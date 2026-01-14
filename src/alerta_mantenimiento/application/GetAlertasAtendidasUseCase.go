//GetAlertasAtendidasUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
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