package entities

import "time"

type Ruta struct {
	RutaID      int32     `json:"ruta_id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	JsonRuta    string    `json:"json_ruta"`
	Eliminado   bool      `json:"eliminado"`
	CreatedAt   time.Time `json:"created_at"`
	// ConductorID es el id_chofer con asignación activa (historial_asignacion_camion)
	// del camión actualmente asignado a la ruta (ruta_camion). Solo lo llena GetActivas.
	ConductorID *int32 `json:"conductor_id"`
}
