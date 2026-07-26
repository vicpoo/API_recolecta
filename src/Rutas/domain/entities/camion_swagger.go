package entities

// swagger:model CreateCamionRequest
type CreateCamionRequest struct {
	Placa            string `json:"placa" binding:"required"`
	Modelo           string `json:"modelo" binding:"required"`
	TipoCamionID     int32  `json:"tipo_camion_id" binding:"required"`
	EsRentado        bool   `json:"es_rentado"`
	DisponibilidadID int32  `json:"disponibilidad_id" binding:"required"`
}

// swagger:model UpdateCamionRequest
type UpdateCamionRequest struct {
	Placa            string `json:"placa"`
	Modelo           string `json:"modelo"`
	TipoCamionID     int32  `json:"tipo_camion_id"`
	EsRentado        bool   `json:"es_rentado"`
	DisponibilidadID int32  `json:"disponibilidad_id"`
}

// swagger:model CamionResponse
type CamionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Camion `json:"data"`
	Code    int    `json:"code"`
}

// swagger:model CamionDetailResponse
type CamionDetailResponse struct {
	Success bool   `json:"success"`
	Data    Camion `json:"data"`
}

// swagger:model CamionListResponse
type CamionListResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Camion `json:"data"`
	Code    int      `json:"code"`
}

// swagger:model CamionListSimpleResponse
type CamionListSimpleResponse struct {
	Success bool     `json:"success"`
	Data    []Camion `json:"data"`
}

// swagger:model CamionMessageResponse
type CamionMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

// swagger:model CamionErrorResponse
type CamionErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}
