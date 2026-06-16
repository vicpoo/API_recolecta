// reporte_conductor.go
package entities

import (
	"time"
)

type ReporteConductor struct {
	ReporteID   int32     `json:"reporte_id" gorm:"column:reporte_id;primaryKey;autoIncrement"`
	ConductorID int32     `json:"conductor_id" gorm:"column:conductor_id;not null"`
	CamionID    int32     `json:"camion_id" gorm:"column:camion_id;not null"`
	RutaID      int32     `json:"ruta_id" gorm:"column:ruta_id;not null"`
	Descripcion string    `json:"descripcion" gorm:"column:descripcion;type:varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
}

// Setters
func (r *ReporteConductor) SetReporteID(reporteID int32) {
	r.ReporteID = reporteID
}

func (r *ReporteConductor) SetConductorID(conductorID int32) {
	r.ConductorID = conductorID
}

func (r *ReporteConductor) SetCamionID(camionID int32) {
	r.CamionID = camionID
}

func (r *ReporteConductor) SetRutaID(rutaID int32) {
	r.RutaID = rutaID
}

func (r *ReporteConductor) SetDescripcion(descripcion string) {
	r.Descripcion = descripcion
}

func (r *ReporteConductor) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

// Getters
func (r *ReporteConductor) GetReporteID() int32 {
	return r.ReporteID
}

func (r *ReporteConductor) GetConductorID() int32 {
	return r.ConductorID
}

func (r *ReporteConductor) GetCamionID() int32 {
	return r.CamionID
}

func (r *ReporteConductor) GetRutaID() int32 {
	return r.RutaID
}

func (r *ReporteConductor) GetDescripcion() string {
	return r.Descripcion
}

func (r *ReporteConductor) GetCreatedAt() time.Time {
	return r.CreatedAt
}

// Constructor básico (para creación)
func NewReporteConductor(conductorID int32, camionID int32, rutaID int32, descripcion string) *ReporteConductor {
	return &ReporteConductor{
		ConductorID: conductorID,
		CamionID:    camionID,
		RutaID:      rutaID,
		Descripcion: descripcion,
		CreatedAt:   time.Now(),
	}
}

// Constructor completo (para actualización o consulta)
func NewReporteConductorCompleto(reporteID int32, conductorID int32, camionID int32, rutaID int32, descripcion string, createdAt time.Time) *ReporteConductor {
	return &ReporteConductor{
		ReporteID:   reporteID,
		ConductorID: conductorID,
		CamionID:    camionID,
		RutaID:      rutaID,
		Descripcion: descripcion,
		CreatedAt:   createdAt,
	}
}

// Método para actualizar descripción
func (r *ReporteConductor) ActualizarDescripcion(descripcion string) {
	r.Descripcion = descripcion
}

// Método para agregar más información a la descripción
func (r *ReporteConductor) AgregarDescripcion(descripcionAdicional string) {
	if r.Descripcion == "" {
		r.Descripcion = descripcionAdicional
	} else {
		r.Descripcion = r.Descripcion + ". " + descripcionAdicional
	}
}

// Método para verificar si el reporte tiene todos los datos requeridos
func (r *ReporteConductor) EsValido() bool {
	return r.ConductorID > 0 && r.CamionID > 0 && r.RutaID > 0 && r.Descripcion != ""
}



func NewReporteConductorParaActualizacion(reporteID int32, conductorID int32, camionID int32, rutaID int32, descripcion string) *ReporteConductor {
	return &ReporteConductor{
		ReporteID:   reporteID,
		ConductorID: conductorID,
		CamionID:    camionID,
		RutaID:      rutaID,
		Descripcion: descripcion,
	}
}