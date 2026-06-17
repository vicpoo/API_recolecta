package entities
import "time"


// swagger:model CreateRegistroMantenimientoRequest
type CreateRegistroMantenimientoRequest struct {
	AlertaID *int32 `json:"alerta_id"`
	CamionID int32 `json:"camion_id" binding:"required"`
	CoordinadorID int32 `json:"coordinador_id"`
	MecanicoResponsable string `json:"mecanico_responsable"`
	FechaRealizada time.Time `json:"fecha_realizada"`
	KilometrajeMantenimiento float64 `json:"kilometraje_mantenimiento"`
	Observaciones string `json:"observaciones"`
}

// swagger:model UpdateRegistroMantenimientoRequest
type UpdateRegistroMantenimientoRequest struct {
	AlertaID *int32 `json:"alerta_id"`
	CamionID int32 `json:"camion_id"`
	CoordinadorID int32 `json:"coordinador_id"`
	MecanicoResponsable string `json:"mecanico_responsable"`
	FechaRealizada time.Time `json:"fecha_realizada"`
	KilometrajeMantenimiento float64 `json:"kilometraje_mantenimiento"`
	Observaciones string `json:"observaciones"`
}

// swagger:model RegistroMantenimientoResponse
type RegistroMantenimientoResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    RegistroMantenimiento `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model RegistroMantenimientoDetailResponse
type RegistroMantenimientoDetailResponse struct {
	Success bool         `json:"success"`
	Data    RegistroMantenimiento `json:"data"`
}

// swagger:model RegistroMantenimientoListResponse
type RegistroMantenimientoListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []RegistroMantenimiento `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model RegistroMantenimientoListSimpleResponse
type RegistroMantenimientoListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []RegistroMantenimiento `json:"data"`
}

// swagger:model RegistroMantenimientoMessageResponse
type RegistroMantenimientoMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
