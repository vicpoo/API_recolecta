package entities

// swagger:model CreateRutaRequest
type CreateRutaRequest struct {
	Nombre      string      `json:"nombre" binding:"required"`
	Descripcion string      `json:"descripcion"`
	JsonRuta    interface{} `json:"json_ruta" binding:"required"`
}

// swagger:model RutaResponse
type RutaResponse struct {
	Success bool `json:"success"`
	Data    Ruta `json:"data"`
}

// RutaActivaItem es el item que devuelve GET /api/rutas/activas.
// Incluye conductor_id resuelto via ruta_camion + historial_asignacion_camion
// (null si la ruta no tiene camión/chofer asignado).
//
// swagger:model RutaActivaItem
type RutaActivaItem struct {
	RutaID      int32       `json:"ruta_id" example:"1"`
	Nombre      string      `json:"nombre" example:"Ruta Centro"`
	Descripcion string      `json:"descripcion" example:"Recolección zona centro"`
	// Geometría ya parseada (array/objeto JSON), no el string crudo de BD.
	JsonRuta    interface{} `json:"json_ruta"`
	Eliminado   bool        `json:"eliminado" example:"false"`
	CreatedAt   string      `json:"created_at" example:"2026-07-22T19:30:00Z"`
	// ID del empleado chofer con asignación activa (fecha_baja IS NULL)
	// del camión más reciente de la ruta. Null si no hay asignación.
	ConductorID *int32 `json:"conductor_id" example:"12"`
}

// swagger:model RutaActivasListResponse
type RutaActivasListResponse struct {
	Success bool            `json:"success" example:"true"`
	Data    []RutaActivaItem `json:"data"`
}
