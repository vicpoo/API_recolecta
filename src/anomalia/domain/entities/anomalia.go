// anomalia.go
package entities

import (
	"time"
)

type Anomalia struct {
	AnomaliaID      int32      `json:"anomalia_id" gorm:"column:anomalia_id;primaryKey;autoIncrement"`
	PuntoID         *int32     `json:"punto_id" gorm:"column:punto_id"`
	TipoAnomalia    string     `json:"tipo_anomalia" gorm:"column:tipo_anomalia;type:varchar(50);not null"`
	Descripcion     string     `json:"descripcion" gorm:"column:descripcion;type:text;not null"`
	FechaReporte    time.Time  `json:"fecha_reporte" gorm:"column:fecha_reporte;not null"`
	Estado          string     `json:"estado" gorm:"column:estado;type:varchar(30);not null"`
	FechaResolucion *time.Time `json:"fecha_resolucion" gorm:"column:fecha_resolucion"`
	IDChoferID      int32      `json:"id_chofer_id" gorm:"column:id_chofer_id;not null"`
}

// Setters
func (a *Anomalia) SetAnomaliaID(anomaliaID int32) {
	a.AnomaliaID = anomaliaID
}

func (a *Anomalia) SetPuntoID(puntoID *int32) {
	a.PuntoID = puntoID
}

func (a *Anomalia) SetTipoAnomalia(tipoAnomalia string) {
	a.TipoAnomalia = tipoAnomalia
}

func (a *Anomalia) SetDescripcion(descripcion string) {
	a.Descripcion = descripcion
}

func (a *Anomalia) SetFechaReporte(fechaReporte time.Time) {
	a.FechaReporte = fechaReporte
}

func (a *Anomalia) SetEstado(estado string) {
	a.Estado = estado
}

func (a *Anomalia) SetFechaResolucion(fechaResolucion *time.Time) {
	a.FechaResolucion = fechaResolucion
}

func (a *Anomalia) SetIDChoferID(idChoferID int32) {
	a.IDChoferID = idChoferID
}

// Getters
func (a *Anomalia) GetAnomaliaID() int32 {
	return a.AnomaliaID
}

func (a *Anomalia) GetPuntoID() *int32 {
	return a.PuntoID
}

func (a *Anomalia) GetTipoAnomalia() string {
	return a.TipoAnomalia
}

func (a *Anomalia) GetDescripcion() string {
	return a.Descripcion
}

func (a *Anomalia) GetFechaReporte() time.Time {
	return a.FechaReporte
}

func (a *Anomalia) GetEstado() string {
	return a.Estado
}

func (a *Anomalia) GetFechaResolucion() *time.Time {
	return a.FechaResolucion
}

func (a *Anomalia) GetIDChoferID() int32 {
	return a.IDChoferID
}

// Constructor básico (para creación)
func NewAnomalia(tipoAnomalia, descripcion string, fechaReporte time.Time, estado string, idChoferID int32) *Anomalia {
	return &Anomalia{
		TipoAnomalia: tipoAnomalia,
		Descripcion:  descripcion,
		FechaReporte: fechaReporte,
		Estado:       estado,
		IDChoferID:   idChoferID,
	}
}

// Constructor con punto (para anomalías específicas de ubicación)
func NewAnomaliaConPunto(puntoID *int32, tipoAnomalia, descripcion string, fechaReporte time.Time, estado string, idChoferID int32) *Anomalia {
	return &Anomalia{
		PuntoID:      puntoID,
		TipoAnomalia: tipoAnomalia,
		Descripcion:  descripcion,
		FechaReporte: fechaReporte,
		Estado:       estado,
		IDChoferID:   idChoferID,
	}
}

// Constructor completo (para actualización o consulta)
func NewAnomaliaCompleta(anomaliaID int32, puntoID *int32, tipoAnomalia, descripcion string, fechaReporte time.Time, estado string, fechaResolucion *time.Time, idChoferID int32) *Anomalia {
	return &Anomalia{
		AnomaliaID:      anomaliaID,
		PuntoID:         puntoID,
		TipoAnomalia:    tipoAnomalia,
		Descripcion:     descripcion,
		FechaReporte:    fechaReporte,
		Estado:          estado,
		FechaResolucion: fechaResolucion,
		IDChoferID:      idChoferID,
	}
}

// Método para marcar como resuelta
func (a *Anomalia) MarcarResuelta(fechaResolucion time.Time) {
	a.Estado = "RESUELTA"
	a.FechaResolucion = &fechaResolucion
}

// Método para marcar como pendiente
func (a *Anomalia) MarcarPendiente() {
	a.Estado = "PENDIENTE"
	a.FechaResolucion = nil
}

// Método para marcar como en proceso
func (a *Anomalia) MarcarEnProceso() {
	a.Estado = "EN_PROCESO"
	a.FechaResolucion = nil
}

// Método para verificar si está resuelta
func (a *Anomalia) EstaResuelta() bool {
	return a.Estado == "RESUELTA"
}

// Método para verificar si está pendiente
func (a *Anomalia) EstaPendiente() bool {
	return a.Estado == "PENDIENTE"
}

// Método para verificar si está en proceso
func (a *Anomalia) EstaEnProceso() bool {
	return a.Estado == "EN_PROCESO"
}

// Método para agregar más información a la descripción
func (a *Anomalia) AgregarDescripcion(descripcionAdicional string) {
	if a.Descripcion == "" {
		a.Descripcion = descripcionAdicional
	} else {
		a.Descripcion = a.Descripcion + "\n" + descripcionAdicional
	}
}

// Método para actualizar tipo de anomalía
func (a *Anomalia) ActualizarTipoAnomalia(nuevoTipo string) {
	a.TipoAnomalia = nuevoTipo
}

// Método para verificar si tiene punto asociado
func (a *Anomalia) TienePunto() bool {
	return a.PuntoID != nil
}

// Método para verificar si tiene fecha de resolución
func (a *Anomalia) TieneFechaResolucion() bool {
	return a.FechaResolucion != nil
}

// Método para obtener una descripción resumida (primeros 100 caracteres)
func (a *Anomalia) GetDescripcionResumida() string {
	if len(a.Descripcion) <= 100 {
		return a.Descripcion
	}
	return a.Descripcion[:100] + "..."
}

// Método para obtener fecha de reporte formateada
func (a *Anomalia) GetFechaReporteFormateada() string {
	return a.FechaReporte.Format("02/01/2006 15:04:05")
}

// Método para obtener fecha de resolución formateada (si existe)
func (a *Anomalia) GetFechaResolucionFormateada() string {
	if a.FechaResolucion != nil {
		return a.FechaResolucion.Format("02/01/2006 15:04:05")
	}
	return ""
}

// TableName especifica el nombre de la tabla para GORM
func (Anomalia) TableName() string {
	return "anomalia"
}