package entities
import "time"


// swagger:model CreateHistorialAsignacionCamionRequest
type CreateHistorialAsignacionCamionRequest struct {
	IDHistorial int `json:"id_historial"`
	IDChofer *int `json:"id_chofer"`
	IDCamion *int `json:"id_camion"`
	FechaAsignacion *time.Time `json:"fecha_asignacion"`
	FechaBaja *time.Time `json:"fecha_baja"`
}

// swagger:model UpdateHistorialAsignacionCamionRequest
type UpdateHistorialAsignacionCamionRequest struct {
	IDHistorial int `json:"id_historial"`
	IDChofer *int `json:"id_chofer"`
	IDCamion *int `json:"id_camion"`
	FechaAsignacion *time.Time `json:"fecha_asignacion"`
	FechaBaja *time.Time `json:"fecha_baja"`
}

// swagger:model HistorialAsignacionCamionResponse
type HistorialAsignacionCamionResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    HistorialAsignacionCamion `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model HistorialAsignacionCamionDetailResponse
type HistorialAsignacionCamionDetailResponse struct {
	Success bool         `json:"success"`
	Data    HistorialAsignacionCamion `json:"data"`
}

// swagger:model HistorialAsignacionCamionListResponse
type HistorialAsignacionCamionListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []HistorialAsignacionCamion `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model HistorialAsignacionCamionListSimpleResponse
type HistorialAsignacionCamionListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []HistorialAsignacionCamion `json:"data"`
}

// swagger:model HistorialAsignacionCamionMessageResponse
type HistorialAsignacionCamionMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
