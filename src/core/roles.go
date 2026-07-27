package core

const (
	// CIUDADANO: role_id del JWT al login de ciudadano (login_ciudadano.go).
	// No es fila de empleado; se alinea con la app móvil (RolUsuario.ciudadano = 5)
	// y con el WS (acepta role_id 5). Evitar 0: en JS `!role_id` lo trata como
	// ausente y SecureVault/móvil también descartan 0.
	ADMIN       = 1
	COORDINADOR = 2
	SUPERVISOR  = 3
	CONDUCTOR   = 4
	CIUDADANO   = 5
)

