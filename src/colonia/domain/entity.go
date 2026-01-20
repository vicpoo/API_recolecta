package domain

import "time"

type Colonia struct {
	ColoniaID int       `json:"colonia_id"`
	Nombre    string    `json:"nombre"`
	Zona      string    `json:"zona"`
	CreatedAt time.Time `json:"created_at"`
}
