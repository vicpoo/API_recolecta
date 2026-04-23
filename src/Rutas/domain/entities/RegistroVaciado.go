package entities

import "time"

type RegistroVaciado struct {
	VaciadoID     int32     `json:"vaciado_id"`
	RellenoID     int32     `json:"relleno_id"`
	RutaCamionID  int32     `json:"ruta_camion_id"`
	Hora          time.Time `json:"hora"`
}
