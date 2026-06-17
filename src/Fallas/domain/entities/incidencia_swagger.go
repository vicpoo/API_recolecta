package entities
import "time"


// swagger:model CreateIncidenciaRequest
type CreateIncidenciaRequest struct {
	PuntoRecoleccionID *int32 `json:"punto_recoleccion_id"`
	ConductorID int32 `json:"conductor_id" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	JsonRuta string `json:"json_ruta"`
	FechaReporte time.Time `json:"fecha_reporte"`
}

// swagger:model UpdateIncidenciaRequest
type UpdateIncidenciaRequest struct {
	PuntoRecoleccionID *int32 `json:"punto_recoleccion_id"`
	ConductorID int32 `json:"conductor_id"`
	Descripcion string `json:"descripcion"`
	JsonRuta string `json:"json_ruta"`
	FechaReporte time.Time `json:"fecha_reporte"`
}

// swagger:model IncidenciaResponse
type IncidenciaResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    Incidencia `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model IncidenciaDetailResponse
type IncidenciaDetailResponse struct {
	Success bool         `json:"success"`
	Data    Incidencia `json:"data"`
}

// swagger:model IncidenciaListResponse
type IncidenciaListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []Incidencia `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model IncidenciaListSimpleResponse
type IncidenciaListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []Incidencia `json:"data"`
}

// swagger:model IncidenciaMessageResponse
type IncidenciaMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
