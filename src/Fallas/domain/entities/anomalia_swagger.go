package entities
import "time"

// swagger:model CreateAnomaliaRequest
type CreateAnomaliaRequest struct {
	PuntoID      *int32    `json:"punto_id"`
	TipoAnomalia string    `json:"tipo_anomalia" binding:"required"`
	Descripcion  string    `json:"descripcion" binding:"required"`
	FechaReporte time.Time `json:"fecha_reporte"`
	Estado       string    `json:"estado"`
	IDChoferID   int32     `json:"id_chofer_id" binding:"required"`
}

// swagger:model UpdateAnomaliaRequest
type UpdateAnomaliaRequest struct {
	PuntoID         *int32     `json:"punto_id"`
	TipoAnomalia    string     `json:"tipo_anomalia"`
	Descripcion     string     `json:"descripcion"`
	FechaReporte    time.Time  `json:"fecha_reporte"`
	Estado          string     `json:"estado"`
	FechaResolucion *time.Time `json:"fecha_resolucion"`
	IDChoferID      int32      `json:"id_chofer_id"`
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
