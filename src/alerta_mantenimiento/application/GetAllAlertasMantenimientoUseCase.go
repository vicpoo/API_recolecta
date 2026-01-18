//GetAllAlertasMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type GetAllAlertasMantenimientoUseCase struct {
	repo repositories.IAlertaMantenimiento
}

func NewGetAllAlertasMantenimientoUseCase(repo repositories.IAlertaMantenimiento) *GetAllAlertasMantenimientoUseCase {
	return &GetAllAlertasMantenimientoUseCase{repo: repo}
}

func (uc *GetAllAlertasMantenimientoUseCase) Run() ([]entities.AlertaMantenimiento, error) {
	return uc.repo.GetAll()
}