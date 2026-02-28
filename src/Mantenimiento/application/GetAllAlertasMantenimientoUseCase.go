//GetAllAlertasMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/Mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
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