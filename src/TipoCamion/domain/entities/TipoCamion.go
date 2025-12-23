package entities

import "time"

type TipoCamion struct {
	TipoCamionID int32 `json:"tipo_camion_id"`
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	CreatedAt time.Time `json:"created_at"`
}