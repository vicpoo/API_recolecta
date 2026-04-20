package application

import (
	"time"
	
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type NotificarTodosUsuariosUseCase struct {
	repo repositories.INotificacion
}

func NewNotificarTodosUsuariosUseCase(repo repositories.INotificacion) *NotificarTodosUsuariosUseCase {
	return &NotificarTodosUsuariosUseCase{repo: repo}
}

type NotificarTodosRequest struct {
	Tipo            string  `json:"tipo" binding:"required"`
	Titulo          string  `json:"titulo" binding:"required"`
	Mensaje         string  `json:"mensaje" binding:"required"`
	CreadoPor       *int32  `json:"creado_por"`
	CamionID        *int32  `json:"camion_id,omitempty"`
	FallaID         *int32  `json:"falla_id,omitempty"`
	MantenimientoID *int32  `json:"mantenimiento_id,omitempty"`
}

func (uc *NotificarTodosUsuariosUseCase) Run(req NotificarTodosRequest) error {
	notificacion := &entities.Notificacion{
		Tipo:                       req.Tipo,
		Titulo:                     req.Titulo,
		Mensaje:                    req.Mensaje,
		Activa:                     true,
		IDCamionRelacionado:        req.CamionID,
		IDFallaRelacionado:         req.FallaID,
		IDMantenimientoRelacionado: req.MantenimientoID,
		CreadoPor:                  req.CreadoPor,
		CreatedAt:                  time.Now(),
	}
	
	return uc.repo.SaveForAllUsers(notificacion)
}