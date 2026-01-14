//GetNotificacionesGlobalesUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesGlobalesUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesGlobalesUseCase(repo repositories.INotificacion) *GetNotificacionesGlobalesUseCase {
	return &GetNotificacionesGlobalesUseCase{repo: repo}
}

func (uc *GetNotificacionesGlobalesUseCase) Run() ([]entities.Notificacion, error) {
	return uc.repo.GetGlobales()
}