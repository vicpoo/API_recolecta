package entities

import "time"

type HistorialAsignacionCamion struct {
	IDHistorial     int        `json:"id_historial"`
	IDChofer        *int       `json:"id_chofer"`    
	IDCamion        *int       `json:"id_camion"`     
	FechaAsignacion *time.Time `json:"fecha_asignacion"`
	FechaBaja       *time.Time `json:"fecha_baja"`
	Eliminado       bool       `json:"eliminado"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
