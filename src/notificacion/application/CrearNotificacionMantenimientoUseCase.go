//CrearNotificacionMantenimientoUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type CrearNotificacionMantenimientoUseCase struct {
	repo repositories.INotificacion
}

func NewCrearNotificacionMantenimientoUseCase(repo repositories.INotificacion) *CrearNotificacionMantenimientoUseCase {
	return &CrearNotificacionMantenimientoUseCase{repo: repo}
}

func (uc *CrearNotificacionMantenimientoUseCase) Run(usuarioID *int32, titulo string, mensaje string, camionID int32, mantenimientoID int32, creadoPor *int32) error {
	return uc.repo.CrearNotificacionMantenimiento(usuarioID, titulo, mensaje, camionID, mantenimientoID, creadoPor)
}