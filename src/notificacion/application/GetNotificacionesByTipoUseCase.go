//GetNotificacionesByTipoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByTipoUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByTipoUseCase(repo repositories.INotificacion) *GetNotificacionesByTipoUseCase {
	return &GetNotificacionesByTipoUseCase{repo: repo}
}

func (uc *GetNotificacionesByTipoUseCase) Run(tipo string) ([]entities.Notificacion, error) {
	return uc.repo.GetByTipo(tipo)
}