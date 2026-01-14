package entities

type RellenoSanitario struct {
	RellenoID        int     `json:"relleno_id"`
	Nombre           string  `json:"nombre"`
	Direccion        string  `json:"direccion"`
	EsRentado        bool    `json:"es_rentado"`
	Eliminado        bool    `json:"eliminado"`
	CapacidadToneladas float64 `json:"capacidad_toneladas"`
}
