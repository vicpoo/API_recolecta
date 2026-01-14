//CrearNotificacionEmergenciaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type CrearNotificacionEmergenciaUseCase struct {
	repo repositories.INotificacion
}

func NewCrearNotificacionEmergenciaUseCase(repo repositories.INotificacion) *CrearNotificacionEmergenciaUseCase {
	return &CrearNotificacionEmergenciaUseCase{repo: repo}
}

func (uc *CrearNotificacionEmergenciaUseCase) Run(usuarioID *int32, titulo string, mensaje string, camionID int32, creadoPor *int32) error {
	return uc.repo.CrearNotificacionEmergencia(usuarioID, titulo, mensaje, camionID, creadoPor)
}