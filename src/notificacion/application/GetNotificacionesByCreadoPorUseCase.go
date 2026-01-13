//GetNotificacionesByCreadoPorUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesByCreadoPorUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesByCreadoPorUseCase(repo repositories.INotificacion) *GetNotificacionesByCreadoPorUseCase {
	return &GetNotificacionesByCreadoPorUseCase{repo: repo}
}

func (uc *GetNotificacionesByCreadoPorUseCase) Run(creadoPor int32) ([]entities.Notificacion, error) {
	return uc.repo.GetByCreadoPor(creadoPor)
}