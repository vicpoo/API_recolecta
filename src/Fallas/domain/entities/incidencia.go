// incidencia.go
package entities

import (
	"time"
)

type Incidencia struct {
	IncidenciaID        int32      `json:"incidencia_id" gorm:"column:incidencia_id;primaryKey;autoIncrement"`
	PuntoRecoleccionID  *int32     `json:"punto_recoleccion_id" gorm:"column:punto_recoleccion_id"`
	ConductorID         int32      `json:"conductor_id" gorm:"column:conductor_id;not null"`
	Descripcion         string     `json:"descripcion" gorm:"column:descripcion;type:varchar(255);not null"`
	JsonRuta            string     `json:"json_ruta" gorm:"column:json_ruta;type:text"`
	FechaReporte        time.Time  `json:"fecha_reporte" gorm:"column:fecha_reporte"`
	Eliminado           bool       `json:"eliminado" gorm:"column:eliminado;default:false"`
	CreatedAt           time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

// Setters
func (i *Incidencia) SetIncidenciaID(incidenciaID int32) {
	i.IncidenciaID = incidenciaID
}

func (i *Incidencia) SetPuntoRecoleccionID(puntoRecoleccionID *int32) {
	i.PuntoRecoleccionID = puntoRecoleccionID
}

func (i *Incidencia) SetConductorID(conductorID int32) {
	i.ConductorID = conductorID
}

func (i *Incidencia) SetDescripcion(descripcion string) {
	i.Descripcion = descripcion
}

func (i *Incidencia) SetJsonRuta(jsonRuta string) {
	i.JsonRuta = jsonRuta
}

func (i *Incidencia) SetFechaReporte(fechaReporte time.Time) {
	i.FechaReporte = fechaReporte
}

func (i *Incidencia) SetEliminado(eliminado bool) {
	i.Eliminado = eliminado
}

func (i *Incidencia) SetCreatedAt(createdAt time.Time) {
	i.CreatedAt = createdAt
}

func (i *Incidencia) SetUpdatedAt(updatedAt time.Time) {
	i.UpdatedAt = updatedAt
}

// Getters
func (i *Incidencia) GetIncidenciaID() int32 {
	return i.IncidenciaID
}

func (i *Incidencia) GetPuntoRecoleccionID() *int32 {
	return i.PuntoRecoleccionID
}

func (i *Incidencia) GetConductorID() int32 {
	return i.ConductorID
}

func (i *Incidencia) GetDescripcion() string {
	return i.Descripcion
}

func (i *Incidencia) GetJsonRuta() string {
	return i.JsonRuta
}

func (i *Incidencia) GetFechaReporte() time.Time {
	return i.FechaReporte
}

func (i *Incidencia) GetEliminado() bool {
	return i.Eliminado
}

func (i *Incidencia) GetCreatedAt() time.Time {
	return i.CreatedAt
}

func (i *Incidencia) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}

// Constructor básico (para creación)
func NewIncidencia(conductorID int32, descripcion string, fechaReporte time.Time) *Incidencia {
	return &Incidencia{
		ConductorID:  conductorID,
		Descripcion:  descripcion,
		FechaReporte: fechaReporte,
		Eliminado:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// Constructor con punto de recolección (para incidencias específicas de ubicación)
func NewIncidenciaConPunto(puntoRecoleccionID *int32, conductorID int32, descripcion string, fechaReporte time.Time) *Incidencia {
	return &Incidencia{
		PuntoRecoleccionID: puntoRecoleccionID,
		ConductorID:        conductorID,
		Descripcion:        descripcion,
		FechaReporte:       fechaReporte,
		Eliminado:          false,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

// Constructor con JSON de ruta (para incidencias con información de ruta)
func NewIncidenciaConRuta(puntoRecoleccionID *int32, conductorID int32, descripcion string, jsonRuta string, fechaReporte time.Time) *Incidencia {
	return &Incidencia{
		PuntoRecoleccionID: puntoRecoleccionID,
		ConductorID:        conductorID,
		Descripcion:        descripcion,
		JsonRuta:           jsonRuta,
		FechaReporte:       fechaReporte,
		Eliminado:          false,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

// Constructor completo (para actualización o consulta)
func NewIncidenciaCompleta(incidenciaID int32, puntoRecoleccionID *int32, conductorID int32, descripcion string, jsonRuta string, fechaReporte time.Time, eliminado bool, createdAt time.Time, updatedAt time.Time) *Incidencia {
	return &Incidencia{
		IncidenciaID:       incidenciaID,
		PuntoRecoleccionID: puntoRecoleccionID,
		ConductorID:        conductorID,
		Descripcion:        descripcion,
		JsonRuta:           jsonRuta,
		FechaReporte:       fechaReporte,
		Eliminado:          eliminado,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}

// Método para agregar información al JSON de ruta
func (i *Incidencia) AgregarRutaJSON(jsonData string) {
	if i.JsonRuta == "" {
		i.JsonRuta = jsonData
	} else {
		i.JsonRuta = i.JsonRuta + "," + jsonData
	}
	i.UpdatedAt = time.Now()
}

// Método para actualizar descripción
func (i *Incidencia) ActualizarDescripcion(descripcion string) {
	i.Descripcion = descripcion
	i.UpdatedAt = time.Now()
}

// Método para marcar como eliminado (borrado lógico)
func (i *Incidencia) MarcarEliminado() {
	i.Eliminado = true
	i.UpdatedAt = time.Now()
}

// Método para restaurar (deshacer borrado lógico)
func (i *Incidencia) Restaurar() {
	i.Eliminado = false
	i.UpdatedAt = time.Now()
}

// Método para verificar si está eliminado
func (i *Incidencia) EstaEliminado() bool {
	return i.Eliminado
}

// Método para verificar si tiene punto de recolección asociado
func (i *Incidencia) TienePuntoRecoleccion() bool {
	return i.PuntoRecoleccionID != nil
}

// Método para verificar si tiene JSON de ruta
func (i *Incidencia) TieneRutaJSON() bool {
	return i.JsonRuta != ""
}

// Método para actualizar fecha de reporte
func (i *Incidencia) ActualizarFechaReporte(fechaReporte time.Time) {
	i.FechaReporte = fechaReporte
	i.UpdatedAt = time.Now()
}