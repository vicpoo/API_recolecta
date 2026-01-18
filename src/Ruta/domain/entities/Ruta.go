package entities

import "time"

type Ruta struct {
	RutaID      int32     `json:"ruta_id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	JsonRuta    string    `json:"json_ruta"`
	Eliminado   bool      `json:"eliminado"`
	CreatedAt   time.Time `json:"created_at"`
}
