package core

const (
	// CIUDADANO no es una fila real de la tabla de roles de empleado -- es el
	// valor centinela que login_ciudadano.go fija a mano como role_id del JWT
	// cuando quien inicia sesion es un ciudadano (ver GenerateToken en
	// login_ciudadano.go). Se declara aqui para no repetir el "0" mágico en
	// cada lugar que necesita distinguir un JWT de ciudadano de uno de
	// empleado (conductor/staff).
	CIUDADANO   = 0
	ADMIN       = 1
	COORDINADOR = 2
	SUPERVISOR  = 3
	CONDUCTOR   = 4
)

