package domain

// swagger:model CreateColoniaRequest
type CreateColoniaRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Zona   string `json:"zona" binding:"required"`
}

// swagger:model UpdateColoniaRequest
type UpdateColoniaRequest struct {
	Nombre string `json:"nombre"`
	Zona   string `json:"zona"`
}

// swagger:model ColoniaResponse
type ColoniaResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Colonia `json:"data"`
	Code    int     `json:"code"`
}

// swagger:model ColoniaDetailResponse
type ColoniaDetailResponse struct {
	Success bool    `json:"success"`
	Data    Colonia `json:"data"`
}

// swagger:model ColoniaListResponse
type ColoniaListResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Colonia `json:"data"`
	Code    int       `json:"code"`
}

// swagger:model ColoniaMessageResponse
type ColoniaMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
