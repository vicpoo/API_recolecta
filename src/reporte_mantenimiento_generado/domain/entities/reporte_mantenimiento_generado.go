// reporte_mantenimiento_generado.go
package entities

import (
	"time"
)

type ReporteMantenimientoGenerado struct {
	ReporteID      int32     `json:"reporte_id" gorm:"column:reporte_id;primaryKey;autoIncrement"`
	CoordinadorID  int32     `json:"coordinador_id" gorm:"column:coordinador_id;not null"`
	FechaDesde     time.Time `json:"fecha_desde" gorm:"column:fecha_desde;not null"`
	FechaHasta     time.Time `json:"fecha_hasta" gorm:"column:fecha_hasta;not null"`
	Observaciones  string    `json:"observaciones" gorm:"column:observaciones;type:varchar(255)"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
}

// Setters
func (r *ReporteMantenimientoGenerado) SetReporteID(reporteID int32) {
	r.ReporteID = reporteID
}

func (r *ReporteMantenimientoGenerado) SetCoordinadorID(coordinadorID int32) {
	r.CoordinadorID = coordinadorID
}

func (r *ReporteMantenimientoGenerado) SetFechaDesde(fechaDesde time.Time) {
	r.FechaDesde = fechaDesde
}

func (r *ReporteMantenimientoGenerado) SetFechaHasta(fechaHasta time.Time) {
	r.FechaHasta = fechaHasta
}

func (r *ReporteMantenimientoGenerado) SetObservaciones(observaciones string) {
	r.Observaciones = observaciones
}

func (r *ReporteMantenimientoGenerado) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

// Getters
func (r *ReporteMantenimientoGenerado) GetReporteID() int32 {
	return r.ReporteID
}

func (r *ReporteMantenimientoGenerado) GetCoordinadorID() int32 {
	return r.CoordinadorID
}

func (r *ReporteMantenimientoGenerado) GetFechaDesde() time.Time {
	return r.FechaDesde
}

func (r *ReporteMantenimientoGenerado) GetFechaHasta() time.Time {
	return r.FechaHasta
}

func (r *ReporteMantenimientoGenerado) GetObservaciones() string {
	return r.Observaciones
}

func (r *ReporteMantenimientoGenerado) GetCreatedAt() time.Time {
	return r.CreatedAt
}

// Constructor básico
func NewReporteMantenimientoGenerado(coordinadorID int32, fechaDesde, fechaHasta time.Time, observaciones string) *ReporteMantenimientoGenerado {
	return &ReporteMantenimientoGenerado{
		CoordinadorID: coordinadorID,
		FechaDesde:    fechaDesde,
		FechaHasta:    fechaHasta,
		Observaciones: observaciones,
		CreatedAt:     time.Now(),
	}
}

// Constructor completo
func NewReporteMantenimientoGeneradoCompleto(reporteID int32, coordinadorID int32, fechaDesde, fechaHasta time.Time, observaciones string, createdAt time.Time) *ReporteMantenimientoGenerado {
	return &ReporteMantenimientoGenerado{
		ReporteID:     reporteID,
		CoordinadorID: coordinadorID,
		FechaDesde:    fechaDesde,
		FechaHasta:    fechaHasta,
		Observaciones: observaciones,
		CreatedAt:     createdAt,
	}
}

// Constructor para actualizaciones
func NewReporteMantenimientoGeneradoParaActualizacion(reporteID int32, coordinadorID int32, fechaDesde, fechaHasta time.Time, observaciones string) *ReporteMantenimientoGenerado {
	return &ReporteMantenimientoGenerado{
		ReporteID:     reporteID,
		CoordinadorID: coordinadorID,
		FechaDesde:    fechaDesde,
		FechaHasta:    fechaHasta,
		Observaciones: observaciones,
		// CreatedAt no se establece aquí para mantener el valor original
	}
}

// Método para agregar más observaciones
func (r *ReporteMantenimientoGenerado) AgregarObservacion(observacionAdicional string) {
	if r.Observaciones == "" {
		r.Observaciones = observacionAdicional
	} else {
		r.Observaciones = r.Observaciones + "\n\n" + observacionAdicional
	}
}

// Método para obtener fecha formateada
func (r *ReporteMantenimientoGenerado) GetFechaDesdeFormateada() string {
	return r.FechaDesde.Format("02/01/2006 15:04:05")
}

func (r *ReporteMantenimientoGenerado) GetFechaHastaFormateada() string {
	return r.FechaHasta.Format("02/01/2006 15:04:05")
}

func (r *ReporteMantenimientoGenerado) GetCreatedAtFormateada() string {
	return r.CreatedAt.Format("02/01/2006 15:04:05")
}

// Método para verificar si tiene observaciones
func (r *ReporteMantenimientoGenerado) TieneObservaciones() bool {
	return r.Observaciones != ""
}

// Método para obtener una vista previa de las observaciones (primeros 100 caracteres)
func (r *ReporteMantenimientoGenerado) GetObservacionesResumidas() string {
	if len(r.Observaciones) <= 100 {
		return r.Observaciones
	}
	return r.Observaciones[:100] + "..."
}

// Método para obtener el rango de fechas como string
func (r *ReporteMantenimientoGenerado) GetRangoFechas() string {
	return r.GetFechaDesdeFormateada() + " - " + r.GetFechaHastaFormateada()
}

// Método para calcular la duración del periodo en días
func (r *ReporteMantenimientoGenerado) GetDuracionEnDias() int {
	duracion := r.FechaHasta.Sub(r.FechaDesde)
	return int(duracion.Hours() / 24)
}

// Método para verificar si las fechas son válidas (fecha_hasta >= fecha_desde)
func (r *ReporteMantenimientoGenerado) FechasValidas() bool {
	return !r.FechaHasta.Before(r.FechaDesde)
}

// TableName especifica el nombre de la tabla para GORM
func (ReporteMantenimientoGenerado) TableName() string {
	return "reporte_mantenimiento_generado"
}