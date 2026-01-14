//CrearNotificacionFallaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type CrearNotificacionFallaUseCase struct {
	repo repositories.INotificacion
}

func NewCrearNotificacionFallaUseCase(repo repositories.INotificacion) *CrearNotificacionFallaUseCase {
	return &CrearNotificacionFallaUseCase{repo: repo}
}

func (uc *CrearNotificacionFallaUseCase) Run(usuarioID *int32, titulo string, mensaje string, camionID int32, fallaID int32, creadoPor *int32) error {
	// Usar el método específico del repositorio
	return uc.repo.CrearNotificacionFalla(usuarioID, titulo, mensaje, camionID, fallaID, creadoPor)
}