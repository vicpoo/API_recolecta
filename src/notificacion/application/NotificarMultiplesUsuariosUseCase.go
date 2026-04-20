package application

import (
	"time"
	
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type NotificarMultiplesUsuariosUseCase struct {
	repo repositories.INotificacion
}

func NewNotificarMultiplesUsuariosUseCase(repo repositories.INotificacion) *NotificarMultiplesUsuariosUseCase {
	return &NotificarMultiplesUsuariosUseCase{repo: repo}
}

type NotificarMultiplesRequest struct {
	UsuarioIDs      []int32 `json:"usuario_ids" binding:"required"`
	Tipo            string  `json:"tipo" binding:"required"`
	Titulo          string  `json:"titulo" binding:"required"`
	Mensaje         string  `json:"mensaje" binding:"required"`
	CreadoPor       *int32  `json:"creado_por"`
	CamionID        *int32  `json:"camion_id,omitempty"`
	FallaID         *int32  `json:"falla_id,omitempty"`
	MantenimientoID *int32  `json:"mantenimiento_id,omitempty"`
}

func (uc *NotificarMultiplesUsuariosUseCase) Run(req NotificarMultiplesRequest) error {
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
	
	return uc.repo.SaveForMultipleUsers(notificacion, req.UsuarioIDs)
}