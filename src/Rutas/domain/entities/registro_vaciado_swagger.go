package entities
import "time"


// swagger:model CreateRegistroVaciadoRequest
type CreateRegistroVaciadoRequest struct {
	RellenoID int32 `json:"relleno_id"`
	RutaCamionID int32 `json:"ruta_camion_id"`
	Hora time.Time `json:"hora"`
}

// swagger:model UpdateRegistroVaciadoRequest
type UpdateRegistroVaciadoRequest struct {
	RellenoID int32 `json:"relleno_id"`
	RutaCamionID int32 `json:"ruta_camion_id"`
	Hora time.Time `json:"hora"`
}

// swagger:model RegistroVaciadoResponse
type RegistroVaciadoResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    RegistroVaciado `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model RegistroVaciadoDetailResponse
type RegistroVaciadoDetailResponse struct {
	Success bool         `json:"success"`
	Data    RegistroVaciado `json:"data"`
}

// swagger:model RegistroVaciadoListResponse
type RegistroVaciadoListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []RegistroVaciado `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model RegistroVaciadoListSimpleResponse
type RegistroVaciadoListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []RegistroVaciado `json:"data"`
}

// swagger:model RegistroVaciadoMessageResponse
type RegistroVaciadoMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
