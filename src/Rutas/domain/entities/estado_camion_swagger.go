package entities
import "time"


// swagger:model CreateEstadoCamionRequest
type CreateEstadoCamionRequest struct {
	EstadoID int32 `json:"estado_id"`
	CamionID int32 `json:"camion_id" binding:"required"`
	Estado string `json:"estado"`
	Timestamp time.Time `json:"timestamp"`
	Observaciones string `json:"observaciones"`
}

// swagger:model UpdateEstadoCamionRequest
type UpdateEstadoCamionRequest struct {
	EstadoID int32 `json:"estado_id"`
	CamionID int32 `json:"camion_id"`
	Estado string `json:"estado"`
	Timestamp time.Time `json:"timestamp"`
	Observaciones string `json:"observaciones"`
}

// swagger:model EstadoCamionResponse
type EstadoCamionResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    EstadoCamion `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model EstadoCamionDetailResponse
type EstadoCamionDetailResponse struct {
	Success bool         `json:"success"`
	Data    EstadoCamion `json:"data"`
}

// swagger:model EstadoCamionListResponse
type EstadoCamionListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []EstadoCamion `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model EstadoCamionListSimpleResponse
type EstadoCamionListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []EstadoCamion `json:"data"`
}

// swagger:model EstadoCamionMessageResponse
type EstadoCamionMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
