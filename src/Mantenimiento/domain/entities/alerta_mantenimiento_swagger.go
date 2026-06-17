package entities


// swagger:model CreateAlertaMantenimientoRequest
type CreateAlertaMantenimientoRequest struct {
	CamionID            int32  `json:"camion_id" binding:"required"`
	TipoMantenimientoID int32  `json:"tipo_mantenimiento_id" binding:"required"`
	Descripcion         string `json:"descripcion"`
	Observaciones       string `json:"observaciones"`
}

// swagger:model UpdateAlertaMantenimientoRequest
type UpdateAlertaMantenimientoRequest struct {
	CamionID            int32  `json:"camion_id"`
	TipoMantenimientoID int32  `json:"tipo_mantenimiento_id"`
	Descripcion         string `json:"descripcion"`
	Observaciones       string `json:"observaciones"`
	Atendido            bool   `json:"atendido"`
}

// swagger:model AlertaMantenimientoResponse
type AlertaMantenimientoResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    AlertaMantenimiento `json:"data"`
	Code    int                 `json:"code"`
}

// swagger:model AlertaMantenimientoDetailResponse
type AlertaMantenimientoDetailResponse struct {
	Success bool                `json:"success"`
	Data    AlertaMantenimiento `json:"data"`
}

// swagger:model AlertaMantenimientoListResponse
type AlertaMantenimientoListResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    []AlertaMantenimiento `json:"data"`
	Code    int                   `json:"code"`
}

// swagger:model AlertaMantenimientoMessageResponse
type AlertaMantenimientoMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
