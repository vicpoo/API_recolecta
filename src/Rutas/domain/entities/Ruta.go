package entities

import "time"

type Ruta struct {
	RutaID      int32     `json:"ruta_id" example:"1"`
	Nombre      string    `json:"nombre" example:"Ruta Centro"`
	Descripcion string    `json:"descripcion" example:"Recolección zona centro"`
	JsonRuta    string    `json:"json_ruta"`
	Eliminado   bool      `json:"eliminado" example:"false"`
	CreatedAt   time.Time `json:"created_at"`
	// ConductorID es el id_chofer con asignación activa (historial_asignacion_camion)
	// del camión actualmente asignado a la ruta (ruta_camion). Solo lo llena GetActivas.
	// swagger:description ID del conductor asignado (null si no hay asignación activa)
	ConductorID *int32 `json:"conductor_id" example:"12"`
}
