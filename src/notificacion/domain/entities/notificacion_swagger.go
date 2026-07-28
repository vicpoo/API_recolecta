package entities



// swagger:model CreateNotificacionRequest
type CreateNotificacionRequest struct {
	UsuarioID *int32 `json:"usuario_id"`
	Tipo string `json:"tipo"`
	Titulo string `json:"titulo"`
	Mensaje string `json:"mensaje"`
	Activa bool `json:"activa"`
	IDCamionRelacionado *int32 `json:"id_camion_relacionado,omitempty"`
	IDFallaRelacionado *int32 `json:"id_falla_relacionado,omitempty"`
	IDMantenimientoRelacionado *int32 `json:"id_mantenimiento_relacionado,omitempty"`
	CreadoPor *int32 `json:"creado_por"`
}

// swagger:model UpdateNotificacionRequest
type UpdateNotificacionRequest struct {
	UsuarioID *int32 `json:"usuario_id"`
	Tipo string `json:"tipo"`
	Titulo string `json:"titulo"`
	Mensaje string `json:"mensaje"`
	Activa bool `json:"activa"`
	IDCamionRelacionado *int32 `json:"id_camion_relacionado,omitempty"`
	IDFallaRelacionado *int32 `json:"id_falla_relacionado,omitempty"`
	IDMantenimientoRelacionado *int32 `json:"id_mantenimiento_relacionado,omitempty"`
	CreadoPor *int32 `json:"creado_por"`
}

// swagger:model NotificacionResponse
type NotificacionResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    Notificacion `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model NotificacionDetailResponse
type NotificacionDetailResponse struct {
	Success bool         `json:"success"`
	Data    Notificacion `json:"data"`
}

// swagger:model NotificacionListResponse
type NotificacionListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []Notificacion `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model NotificacionListSimpleResponse
type NotificacionListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []Notificacion `json:"data"`
}

// swagger:model NotificacionMessageResponse
type NotificacionMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
