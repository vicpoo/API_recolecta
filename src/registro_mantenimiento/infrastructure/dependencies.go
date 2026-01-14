// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
)

func InitRegistroMantenimientoDependencies() (
	*CreateRegistroMantenimientoController,
	*GetRegistroMantenimientoByIDController,
	*UpdateRegistroMantenimientoController,
	*DeleteRegistroMantenimientoController,
	*GetAllRegistrosMantenimientoController,
	*GetRegistroByAlertaIDController,
	*GetRegistrosByCamionIDController,
	*GetRegistrosByCoordinadorIDController,
	*GetRegistrosByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresRegistroMantenimientoRepository()

	// Casos de uso CRUD básicos
	createUseCase := application.NewCreateRegistroMantenimientoUseCase(repo)
	getByIDUseCase := application.NewGetRegistroMantenimientoByIDUseCase(repo)
	updateUseCase := application.NewUpdateRegistroMantenimientoUseCase(repo)
	deleteUseCase := application.NewDeleteRegistroMantenimientoUseCase(repo)
	getAllUseCase := application.NewGetAllRegistrosMantenimientoUseCase(repo)
	
	// Casos de uso específicos
	getByAlertaIDUseCase := application.NewGetRegistroByAlertaIDUseCase(repo)
	getByCamionIDUseCase := application.NewGetRegistrosByCamionIDUseCase(repo)
	getByCoordinadorIDUseCase := application.NewGetRegistrosByCoordinadorIDUseCase(repo)
	getByFechaRangeUseCase := application.NewGetRegistrosByFechaRangeUseCase(repo)

	// Controladores CRUD básicos
	createController := NewCreateRegistroMantenimientoController(createUseCase)
	getByIDController := NewGetRegistroMantenimientoByIDController(getByIDUseCase)
	updateController := NewUpdateRegistroMantenimientoController(updateUseCase)
	deleteController := NewDeleteRegistroMantenimientoController(deleteUseCase)
	getAllController := NewGetAllRegistrosMantenimientoController(getAllUseCase)
	
	// Controladores específicos
	getByAlertaIDController := NewGetRegistroByAlertaIDController(getByAlertaIDUseCase)
	getByCamionIDController := NewGetRegistrosByCamionIDController(getByCamionIDUseCase)
	getByCoordinadorIDController := NewGetRegistrosByCoordinadorIDController(getByCoordinadorIDUseCase)
	getByFechaRangeController := NewGetRegistrosByFechaRangeController(getByFechaRangeUseCase)

	return createController, 
	       getByIDController, 
	       updateController, 
	       deleteController, 
	       getAllController,
	       getByAlertaIDController,
	       getByCamionIDController,
	       getByCoordinadorIDController,
	       getByFechaRangeController
}