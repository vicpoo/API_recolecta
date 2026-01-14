package entities

import "time"

type PuntoRecoleccion struct {
	PuntoID   int32     `json:"punto_id"`
	RutaID    int32    `json:"ruta_id"`
	CP        string    `json:"cp"`
	Eliminado bool      `json:"eliminado"`
	CreatedAt time.Time `json:"created_at"`
}
