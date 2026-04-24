package entities

import "time"

type Domicilio struct {
	ID          int       `json:"id"`
	CiudadanoID int       `json:"ciudadano_id"`
	ColoniaID   int       `json:"colonia_id"`
	Alias       string    `json:"alias"`
	Calle       string    `json:"calle"`
	Numero      string    `json:"numero"`
	Referencia  *string   `json:"referencia,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}