//GetNotificacionesActivasUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type GetNotificacionesActivasUseCase struct {
	repo repositories.INotificacion
}

func NewGetNotificacionesActivasUseCase(repo repositories.INotificacion) *GetNotificacionesActivasUseCase {
	return &GetNotificacionesActivasUseCase{repo: repo}
}

func (uc *GetNotificacionesActivasUseCase) Run() ([]entities.Notificacion, error) {
	return uc.repo.GetActivas()
}