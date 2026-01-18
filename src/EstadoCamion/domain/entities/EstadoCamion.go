package entities

import "time"

type EstadoCamion struct {
	EstadoID      int32      `json:"estado_id"`
	CamionID      int32     `json:"camion_id"`      
	Estado        string    `json:"estado"`           // nullable
	Timestamp     time.Time `json:"timestamp"`        // nullable
	Observaciones string    `json:"observaciones"`    // nullable
}
