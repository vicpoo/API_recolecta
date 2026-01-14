// registro_mantenimiento.go
package entities

import (
	"time"
)

type RegistroMantenimiento struct {
	RegistroID              int32     `json:"registro_id" gorm:"column:registro_id;primaryKey;autoIncrement"`
	AlertaID                *int32    `json:"alerta_id" gorm:"column:alerta_id"`
	CamionID                int32     `json:"camion_id" gorm:"column:camion_id;not null"`
	CoordinadorID           int32     `json:"coordinador_id" gorm:"column:coordinador_id;not null"`
	MecanicoResponsable     string    `json:"mecanico_responsable" gorm:"column:mecanico_responsable;type:varchar(255);not null"`
	FechaRealizada          time.Time `json:"fecha_realizada" gorm:"column:fecha_realizada"`
	KilometrajeMantenimiento float64   `json:"kilometraje_mantenimiento" gorm:"column:kilometraje_mantenimiento"`
	Observaciones           string    `json:"observaciones" gorm:"column:observaciones;type:text"`
	CreatedAt               time.Time `json:"created_at" gorm:"column:created_at"`
}

// Setters
func (r *RegistroMantenimiento) SetRegistroID(registroID int32) {
	r.RegistroID = registroID
}

func (r *RegistroMantenimiento) SetAlertaID(alertaID *int32) {
	r.AlertaID = alertaID
}

func (r *RegistroMantenimiento) SetCamionID(camionID int32) {
	r.CamionID = camionID
}

func (r *RegistroMantenimiento) SetCoordinadorID(coordinadorID int32) {
	r.CoordinadorID = coordinadorID
}

func (r *RegistroMantenimiento) SetMecanicoResponsable(mecanicoResponsable string) {
	r.MecanicoResponsable = mecanicoResponsable
}

func (r *RegistroMantenimiento) SetFechaRealizada(fechaRealizada time.Time) {
	r.FechaRealizada = fechaRealizada
}

func (r *RegistroMantenimiento) SetKilometrajeMantenimiento(kilometraje float64) {
	r.KilometrajeMantenimiento = kilometraje
}

func (r *RegistroMantenimiento) SetObservaciones(observaciones string) {
	r.Observaciones = observaciones
}

func (r *RegistroMantenimiento) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

// Getters
func (r *RegistroMantenimiento) GetRegistroID() int32 {
	return r.RegistroID
}

func (r *RegistroMantenimiento) GetAlertaID() *int32 {
	return r.AlertaID
}

func (r *RegistroMantenimiento) GetCamionID() int32 {
	return r.CamionID
}

func (r *RegistroMantenimiento) GetCoordinadorID() int32 {
	return r.CoordinadorID
}

func (r *RegistroMantenimiento) GetMecanicoResponsable() string {
	return r.MecanicoResponsable
}

func (r *RegistroMantenimiento) GetFechaRealizada() time.Time {
	return r.FechaRealizada
}

func (r *RegistroMantenimiento) GetKilometrajeMantenimiento() float64 {
	return r.KilometrajeMantenimiento
}

func (r *RegistroMantenimiento) GetObservaciones() string {
	return r.Observaciones
}

func (r *RegistroMantenimiento) GetCreatedAt() time.Time {
	return r.CreatedAt
}

// Constructor básico (para creación)
func NewRegistroMantenimiento(camionID int32, coordinadorID int32, mecanicoResponsable string, fechaRealizada time.Time, kilometraje float64, observaciones string) *RegistroMantenimiento {
	return &RegistroMantenimiento{
		CamionID:                camionID,
		CoordinadorID:           coordinadorID,
		MecanicoResponsable:     mecanicoResponsable,
		FechaRealizada:          fechaRealizada,
		KilometrajeMantenimiento: kilometraje,
		Observaciones:           observaciones,
		CreatedAt:               time.Now(),
	}
}

// Constructor con alerta (para cuando se registra un mantenimiento desde una alerta)
func NewRegistroMantenimientoConAlerta(alertaID *int32, camionID int32, coordinadorID int32, mecanicoResponsable string, fechaRealizada time.Time, kilometraje float64, observaciones string) *RegistroMantenimiento {
	return &RegistroMantenimiento{
		AlertaID:                alertaID,
		CamionID:                camionID,
		CoordinadorID:           coordinadorID,
		MecanicoResponsable:     mecanicoResponsable,
		FechaRealizada:          fechaRealizada,
		KilometrajeMantenimiento: kilometraje,
		Observaciones:           observaciones,
		CreatedAt:               time.Now(),
	}
}

// Constructor completo (para actualización o consulta)
func NewRegistroMantenimientoCompleto(registroID int32, alertaID *int32, camionID int32, coordinadorID int32, mecanicoResponsable string, fechaRealizada time.Time, kilometraje float64, observaciones string, createdAt time.Time) *RegistroMantenimiento {
	return &RegistroMantenimiento{
		RegistroID:              registroID,
		AlertaID:                alertaID,
		CamionID:                camionID,
		CoordinadorID:           coordinadorID,
		MecanicoResponsable:     mecanicoResponsable,
		FechaRealizada:          fechaRealizada,
		KilometrajeMantenimiento: kilometraje,
		Observaciones:           observaciones,
		CreatedAt:               createdAt,
	}
}

// Método para agregar observaciones
func (r *RegistroMantenimiento) AgregarObservacion(observacion string) {
	if r.Observaciones == "" {
		r.Observaciones = observacion
	} else {
		r.Observaciones = r.Observaciones + "\n" + observacion
	}
}

// Método para verificar si tiene alerta asociada
func (r *RegistroMantenimiento) TieneAlertaAsociada() bool {
	return r.AlertaID != nil
}

// Método para actualizar fecha de realización
func (r *RegistroMantenimiento) ActualizarFechaRealizada(fechaRealizada time.Time) {
	r.FechaRealizada = fechaRealizada
}

// Método para actualizar kilometraje
func (r *RegistroMantenimiento) ActualizarKilometraje(kilometraje float64) {
	r.KilometrajeMantenimiento = kilometraje
}