package entities

import "time"

type Camion struct {
	CamionID              int32     `json:"camion_id"`
	Placa                 string    `json:"placa"`
	Modelo                string    `json:"modelo"`
	TipoCamionID          int32     `json:"tipo_camion_id"`
	EsRentado             bool      `json:"es_rentado"`
	DisponibilidadID      int32     `json:"disponibilidad_id"`
	NombreDisponibilidad  string    `json:"nombre_disponibilidad"`
	ColorDisponibilidad   string    `json:"color_disponibilidad"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
