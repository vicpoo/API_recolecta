// alerta_mantenimiento.go
package entities

import (
	"time"
)

type AlertaMantenimiento struct {
	AlertaID            int32     `json:"alerta_id" gorm:"column:alerta_id;primaryKey;autoIncrement"`
	CamionID            int32     `json:"camion_id" gorm:"column:camion_id;not null"`
	TipoMantenimientoID int32     `json:"tipo_mantenimiento_id" gorm:"column:tipo_mantenimiento_id;not null"`
	Descripcion         string    `json:"descripcion" gorm:"column:descripcion;type:varchar(255)"`
	Observaciones       string    `json:"observaciones" gorm:"column:observaciones;type:text"`
	CreatedAt           time.Time `json:"created_at" gorm:"column:created_at"`
	Atendido            bool      `json:"atendido" gorm:"column:atendido;default:false"`
}

// Setters
func (a *AlertaMantenimiento) SetAlertaID(alertaID int32) {
	a.AlertaID = alertaID
}

func (a *AlertaMantenimiento) SetCamionID(camionID int32) {
	a.CamionID = camionID
}

func (a *AlertaMantenimiento) SetTipoMantenimientoID(tipoMantenimientoID int32) {
	a.TipoMantenimientoID = tipoMantenimientoID
}

func (a *AlertaMantenimiento) SetDescripcion(descripcion string) {
	a.Descripcion = descripcion
}

func (a *AlertaMantenimiento) SetObservaciones(observaciones string) {
	a.Observaciones = observaciones
}

func (a *AlertaMantenimiento) SetCreatedAt(createdAt time.Time) {
	a.CreatedAt = createdAt
}

func (a *AlertaMantenimiento) SetAtendido(atendido bool) {
	a.Atendido = atendido
}

// Getters
func (a *AlertaMantenimiento) GetAlertaID() int32 {
	return a.AlertaID
}

func (a *AlertaMantenimiento) GetCamionID() int32 {
	return a.CamionID
}

func (a *AlertaMantenimiento) GetTipoMantenimientoID() int32 {
	return a.TipoMantenimientoID
}

func (a *AlertaMantenimiento) GetDescripcion() string {
	return a.Descripcion
}

func (a *AlertaMantenimiento) GetObservaciones() string {
	return a.Observaciones
}

func (a *AlertaMantenimiento) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *AlertaMantenimiento) GetAtendido() bool {
	return a.Atendido
}

func NewAlertaMantenimiento(camionID int32, tipoMantenimientoID int32, descripcion string, observaciones string) *AlertaMantenimiento {
	return &AlertaMantenimiento{
		CamionID:            camionID,
		TipoMantenimientoID: tipoMantenimientoID,
		Descripcion:         descripcion,
		Observaciones:       observaciones,
		CreatedAt:           time.Now(),
		Atendido:            false,
	}
}

// Constructor con todos los campos (para actualización)
func NewAlertaMantenimientoCompleta(alertaID int32, camionID int32, tipoMantenimientoID int32, descripcion string, observaciones string, createdAt time.Time, atendido bool) *AlertaMantenimiento {
	return &AlertaMantenimiento{
		AlertaID:            alertaID,
		CamionID:            camionID,
		TipoMantenimientoID: tipoMantenimientoID,
		Descripcion:         descripcion,
		Observaciones:       observaciones,
		CreatedAt:           createdAt,
		Atendido:            atendido,
	}
}

// Método para marcar como atendido
func (a *AlertaMantenimiento) MarcarAtendido() {
	a.Atendido = true
}

// Método para actualizar observaciones
func (a *AlertaMantenimiento) AgregarObservacion(observacion string) {
	if a.Observaciones == "" {
		a.Observaciones = observacion
	} else {
		a.Observaciones = a.Observaciones + "\n" + observacion
	}
}

// Método para verificar si está pendiente
func (a *AlertaMantenimiento) EstaPendiente() bool {
	return !a.Atendido
}