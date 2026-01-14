//GetNotificacionesInactivasUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesInactivasUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesInactivasUseCase(repo repositories.INotificacion) *GetNotificacionesInactivasUseCase {
	return &GetNotificacionesInactivasUseCase{repo: repo}
}

func (uc *GetNotificacionesInactivasUseCase) Run() ([]entities.Notificacion, error) {
	return uc.repo.GetInactivas()
}