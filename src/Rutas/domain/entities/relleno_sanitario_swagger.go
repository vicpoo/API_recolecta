package entities

// swagger:model CreateRellenoSanitarioRequest
type CreateRellenoSanitarioRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Direccion string `json:"direccion"`
	EsRentado bool `json:"es_rentado"`
	CapacidadToneladas float64 `json:"capacidad_toneladas"`
}

// swagger:model UpdateRellenoSanitarioRequest
type UpdateRellenoSanitarioRequest struct {
	Nombre string `json:"nombre"`
	Direccion string `json:"direccion"`
	EsRentado bool `json:"es_rentado"`
	CapacidadToneladas float64 `json:"capacidad_toneladas"`
}

// swagger:model RellenoSanitarioResponse
type RellenoSanitarioResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    RellenoSanitario `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model RellenoSanitarioDetailResponse
type RellenoSanitarioDetailResponse struct {
	Success bool         `json:"success"`
	Data    RellenoSanitario `json:"data"`
}

// swagger:model RellenoSanitarioListResponse
type RellenoSanitarioListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []RellenoSanitario `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model RellenoSanitarioListSimpleResponse
type RellenoSanitarioListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []RellenoSanitario `json:"data"`
}

// swagger:model RellenoSanitarioMessageResponse
type RellenoSanitarioMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
