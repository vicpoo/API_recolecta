package entities

import "time"

// swagger:model CreateAnomaliaRequest
type CreateAnomaliaRequest struct {
	TipoAnomalia         TipoAnomalia `json:"tipo_anomalia" binding:"required"` // ANOMALIA | INCIDENCIA | REPORTE_CONDUCTOR | REPORTE_FALLA_CRITICA | SEGUIMIENTO_FALLA_CRITICA
	PuntoID              *int32       `json:"punto_id"`
	ConductorID          *int32       `json:"conductor_id"`
	CamionID             *int32       `json:"camion_id"`
	RutaID               *int32       `json:"ruta_id"`
	AnomaliaReferenciaID *int32       `json:"anomalia_referencia_id"`
	Descripcion          string       `json:"descripcion" binding:"required"`
	JsonRuta             string       `json:"json_ruta"`
	Estado               string       `json:"estado"`
	FechaReporte         time.Time    `json:"fecha_reporte"`
}

// swagger:model UpdateAnomaliaRequest
type UpdateAnomaliaRequest struct {
	TipoAnomalia         TipoAnomalia `json:"tipo_anomalia" binding:"required"`
	PuntoID              *int32       `json:"punto_id"`
	ConductorID          *int32       `json:"conductor_id"`
	CamionID             *int32       `json:"camion_id"`
	RutaID               *int32       `json:"ruta_id"`
	AnomaliaReferenciaID *int32       `json:"anomalia_referencia_id"`
	Descripcion          string       `json:"descripcion" binding:"required"`
	JsonRuta             string       `json:"json_ruta"`
	Estado               string       `json:"estado"`
	Eliminado            bool         `json:"eliminado"`
	FechaReporte         time.Time    `json:"fecha_reporte"`
	FechaResolucion      *time.Time   `json:"fecha_resolucion"`
}

// swagger:model AnomaliaResponse
type AnomaliaResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    Anomalia `json:"data"`
	Code    int      `json:"code"`
}

// swagger:model AnomaliaDetailResponse
type AnomaliaDetailResponse struct {
	Success bool     `json:"success"`
	Data    Anomalia `json:"data"`
}

// swagger:model AnomaliaListResponse
type AnomaliaListResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    []Anomalia `json:"data"`
	Code    int        `json:"code"`
}

// swagger:model AnomaliaMessageResponse
type AnomaliaMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
