//GetNotificacionesByCamionYTipoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByCamionYTipoUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByCamionYTipoUseCase(repo repositories.INotificacion) *GetNotificacionesByCamionYTipoUseCase {
	return &GetNotificacionesByCamionYTipoUseCase{repo: repo}
}

func (uc *GetNotificacionesByCamionYTipoUseCase) Run(camionID int32, tipo string) ([]entities.Notificacion, error) {
	return uc.repo.GetByCamionYTipo(camionID, tipo)
}