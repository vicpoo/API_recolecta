//GetRegistroByAlertaIDUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type GetRegistroByAlertaIDUseCase struct {
	repo repositories.IRegistroMantenimiento
}

func NewGetRegistroByAlertaIDUseCase(repo repositories.IRegistroMantenimiento) *GetRegistroByAlertaIDUseCase {
	return &GetRegistroByAlertaIDUseCase{repo: repo}
}

func (uc *GetRegistroByAlertaIDUseCase) Run(alertaID int32) (*entities.RegistroMantenimiento, error) {
	return uc.repo.GetByAlertaID(alertaID)
}