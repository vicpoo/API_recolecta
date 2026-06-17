package entities



// swagger:model CreateReporteFallaCriticaRequest
type CreateReporteFallaCriticaRequest struct {
	CamionID int32 `json:"camion_id" binding:"required"`
	ConductorID int32 `json:"conductor_id" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

// swagger:model UpdateReporteFallaCriticaRequest
type UpdateReporteFallaCriticaRequest struct {
	CamionID int32 `json:"camion_id"`
	ConductorID int32 `json:"conductor_id"`
	Descripcion string `json:"descripcion"`
}

// swagger:model ReporteFallaCriticaResponse
type ReporteFallaCriticaResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    ReporteFallaCritica `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model ReporteFallaCriticaDetailResponse
type ReporteFallaCriticaDetailResponse struct {
	Success bool         `json:"success"`
	Data    ReporteFallaCritica `json:"data"`
}

// swagger:model ReporteFallaCriticaListResponse
type ReporteFallaCriticaListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []ReporteFallaCritica `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model ReporteFallaCriticaListSimpleResponse
type ReporteFallaCriticaListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []ReporteFallaCritica `json:"data"`
}

// swagger:model ReporteFallaCriticaMessageResponse
type ReporteFallaCriticaMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
