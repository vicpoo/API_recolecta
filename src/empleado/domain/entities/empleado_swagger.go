package entities


// swagger:model CreateEmpleadoRequest
type CreateEmpleadoRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
	Mail string `json:"mail" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Desactivado bool `json:"desactivado"`
	RolID int `json:"rol_id"`
}

// swagger:model UpdateEmpleadoRequest
type UpdateEmpleadoRequest struct {
	Nombre string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Mail string `json:"mail"`
	Username string `json:"username"`
	Password string `json:"password"`
	Desactivado bool `json:"desactivado"`
	RolID int `json:"rol_id"`
}
// swagger:model LoginEmpleadoRequest
type LoginEmpleadoRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// swagger:model LoginEmpleadoResponse
type LoginEmpleadoResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Token   string   `json:"token"`
	Data    Empleado `json:"data"`
	Code    int      `json:"code"`
}


// swagger:model EmpleadoResponse
type EmpleadoResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    Empleado `json:"data"`
	Code    int          `json:"code"`
}

// swagger:model EmpleadoDetailResponse
type EmpleadoDetailResponse struct {
	Success bool         `json:"success"`
	Data    Empleado `json:"data"`
}

// swagger:model EmpleadoListResponse
type EmpleadoListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []Empleado `json:"data"`
	Code    int            `json:"code"`
}

// swagger:model EmpleadoListSimpleResponse
type EmpleadoListSimpleResponse struct {
	Success bool           `json:"success"`
	Data    []Empleado `json:"data"`
}

// swagger:model EmpleadoMessageResponse
type EmpleadoMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
