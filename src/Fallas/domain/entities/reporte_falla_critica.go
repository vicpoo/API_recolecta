// reporte_falla_critica.go
package entities

import (
	"time"
)

type ReporteFallaCritica struct {
	FallaID     int32     `json:"falla_id" gorm:"column:falla_id;primaryKey;autoIncrement"`
	CamionID    int32     `json:"camion_id" gorm:"column:camion_id;not null"`
	ConductorID int32     `json:"conductor_id" gorm:"column:conductor_id;not null"`
	Descripcion string    `json:"descripcion" gorm:"column:descripcion;type:varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	Eliminado   bool      `json:"eliminado" gorm:"column:eliminado;default:false"`
}

// Setters
func (r *ReporteFallaCritica) SetFallaID(fallaID int32) {
	r.FallaID = fallaID
}

func (r *ReporteFallaCritica) SetCamionID(camionID int32) {
	r.CamionID = camionID
}

func (r *ReporteFallaCritica) SetConductorID(conductorID int32) {
	r.ConductorID = conductorID
}

func (r *ReporteFallaCritica) SetDescripcion(descripcion string) {
	r.Descripcion = descripcion
}

func (r *ReporteFallaCritica) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

func (r *ReporteFallaCritica) SetEliminado(eliminado bool) {
	r.Eliminado = eliminado
}

// Getters
func (r *ReporteFallaCritica) GetFallaID() int32 {
	return r.FallaID
}

func (r *ReporteFallaCritica) GetCamionID() int32 {
	return r.CamionID
}

func (r *ReporteFallaCritica) GetConductorID() int32 {
	return r.ConductorID
}

func (r *ReporteFallaCritica) GetDescripcion() string {
	return r.Descripcion
}

func (r *ReporteFallaCritica) GetCreatedAt() time.Time {
	return r.CreatedAt
}

func (r *ReporteFallaCritica) GetEliminado() bool {
	return r.Eliminado
}

// Constructor básico (para creación)
func NewReporteFallaCritica(camionID int32, conductorID int32, descripcion string) *ReporteFallaCritica {
	return &ReporteFallaCritica{
		CamionID:    camionID,
		ConductorID: conductorID,
		Descripcion: descripcion,
		CreatedAt:   time.Now(),
		Eliminado:   false,
	}
}

// Constructor completo (para actualización o consulta)
func NewReporteFallaCriticaCompleto(fallaID int32, camionID int32, conductorID int32, descripcion string, createdAt time.Time, eliminado bool) *ReporteFallaCritica {
	return &ReporteFallaCritica{
		FallaID:     fallaID,
		CamionID:    camionID,
		ConductorID: conductorID,
		Descripcion: descripcion,
		CreatedAt:   createdAt,
		Eliminado:   eliminado,
	}
}

// Método para marcar como eliminado (soft delete)
func (r *ReporteFallaCritica) MarcarComoEliminado() {
	r.Eliminado = true
}

// Método para restaurar (deshacer soft delete)
func (r *ReporteFallaCritica) Restaurar() {
	r.Eliminado = false
}

// Método para verificar si está eliminado
func (r *ReporteFallaCritica) EstaEliminado() bool {
	return r.Eliminado
}

// Método para agregar más información a la descripción
func (r *ReporteFallaCritica) AgregarDescripcion(descripcionAdicional string) {
	if r.Descripcion == "" {
		r.Descripcion = descripcionAdicional
	} else {
		r.Descripcion = r.Descripcion + "\n" + descripcionAdicional
	}
}

// Método para actualizar la descripción
func (r *ReporteFallaCritica) ActualizarDescripcion(nuevaDescripcion string) {
	r.Descripcion = nuevaDescripcion
}

// Método para obtener una descripción resumida (primeros 100 caracteres)
func (r *ReporteFallaCritica) GetDescripcionResumida() string {
	if len(r.Descripcion) <= 100 {
		return r.Descripcion
	}
	return r.Descripcion[:100] + "..."
}

// TableName especifica el nombre de la tabla para GORM
func (ReporteFallaCritica) TableName() string {
	return "reporte_falla_critica"
}