//NotificarUsuarioUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
	"time"
)

type NotificarUsuarioUseCase struct {
	repo repositories.INotificacion
}

func NewNotificarUsuarioUseCase(repo repositories.INotificacion) *NotificarUsuarioUseCase {
	return &NotificarUsuarioUseCase{repo: repo}
}

func (uc *NotificarUsuarioUseCase) Run(creadorID int32, destinatarioID int32, tipo string, titulo string, mensaje string, idCamion *int32, idFalla *int32, idMantenimiento *int32) error {
	notificacion := entities.NewNotificacionCompleta(
		0, // ID será generado
		&destinatarioID,
		tipo,
		titulo,
		mensaje,
		true,
		idCamion,
		idFalla,
		idMantenimiento,
		&creadorID,
		time.Now(),
	)
	
	return uc.repo.Save(notificacion)
}