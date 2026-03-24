package entities

import "time"

type CiudadanoRedis struct {
	UserID    string    `redis:"user_id"`
	FCMToken  string    `redis:"fcm_token"`
	Longitude float64   `redis:"longitude"`
	Latitude  float64   `redis:"latitude"`
	UpdatedAt time.Time `redis:"updated_at"`
}
