package entities

import "time"

type RutaCamion struct {
	RutaCamionID int32     `json:"ruta_camion_id"`
	RutaID       int32     `json:"ruta_id"`
	CamionID     int32     `json:"camion_id"`
	Fecha        time.Time `json:"fecha"`
	CreatedAt    time.Time `json:"created_at"`
	Eliminado    bool      `json:"eliminado"`
}
