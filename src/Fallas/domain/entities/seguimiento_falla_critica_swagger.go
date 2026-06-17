package entities

// swagger:model CreateSeguimientoFallaCriticaRequest
type CreateSeguimientoFallaCriticaRequest struct {
	FallaID    int32  `json:"falla_id" binding:"required"`
	Comentario string `json:"comentario" binding:"required"`
}

// swagger:model UpdateSeguimientoFallaCriticaRequest
type UpdateSeguimientoFallaCriticaRequest struct {
	FallaID    int32  `json:"falla_id"`
	Comentario string `json:"comentario"`
}

// swagger:model SeguimientoFallaCriticaResponse
type SeguimientoFallaCriticaResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Data    SeguimientoFallaCritica `json:"data"`
	Code    int                     `json:"code"`
}

// swagger:model SeguimientoFallaCriticaDetailResponse
type SeguimientoFallaCriticaDetailResponse struct {
	Success bool                    `json:"success"`
	Data    SeguimientoFallaCritica `json:"data"`
}

// swagger:model SeguimientoFallaCriticaListResponse
type SeguimientoFallaCriticaListResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Data    []SeguimientoFallaCritica `json:"data"`
	Code    int                       `json:"code"`
}

// swagger:model SeguimientoFallaCriticaListSimpleResponse
type SeguimientoFallaCriticaListSimpleResponse struct {
	Success bool                      `json:"success"`
	Data    []SeguimientoFallaCritica `json:"data"`
}

// swagger:model SeguimientoFallaCriticaMessageResponse
type SeguimientoFallaCriticaMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
