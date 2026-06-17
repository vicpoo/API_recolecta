package entities



// swagger:model CreateReporteConductorRequest
type CreateReporteConductorRequest struct {
	ConductorID int32 `json:"conductor_id" binding:"required"`
	CamionID int32 `json:"camion_id" binding:"required"`
	RutaID int32 `json:"ruta_id"`
	Descripcion string `json:"descripcion" binding:"required"`
}

// swagger:model UpdateReporteConductorRequest
type UpdateReporteConductorRequest struct {
	ConductorID int32 `json:"conductor_id"`
	CamionID int32 `json:"camion_id"`
	RutaID int32 `json:"ruta_id"`
	Descripcion string `json:"descripcion"`
}

// swagger:model ReporteConductorResponse
type ReporteConductorResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    ReporteConductor `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model ReporteConductorDetailResponse
type ReporteConductorDetailResponse struct {
	Success bool         `json:"success"`
	Data    ReporteConductor `json:"data"`
}

// swagger:model ReporteConductorListResponse
type ReporteConductorListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []ReporteConductor `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model ReporteConductorListSimpleResponse
type ReporteConductorListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []ReporteConductor `json:"data"`
}

// swagger:model ReporteConductorMessageResponse
type ReporteConductorMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
