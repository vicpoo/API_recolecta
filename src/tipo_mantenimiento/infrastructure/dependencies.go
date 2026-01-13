// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/application"
)

func InitTipoMantenimientoDependencies() (
	*CreateTipoMantenimientoController,
	*GetTipoMantenimientoByIDController,
	*UpdateTipoMantenimientoController,
	*DeleteTipoMantenimientoController,
	*GetAllTiposMantenimientoController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresTipoMantenimientoRepository()

	// Casos de uso
	createUseCase := application.NewCreateTipoMantenimientoUseCase(repo)
	getByIDUseCase := application.NewGetTipoMantenimientoByIDUseCase(repo)
	updateUseCase := application.NewUpdateTipoMantenimientoUseCase(repo)
	deleteUseCase := application.NewDeleteTipoMantenimientoUseCase(repo)
	getAllUseCase := application.NewGetAllTiposMantenimientoUseCase(repo)

	// Controladores
	createController := NewCreateTipoMantenimientoController(createUseCase)
	getByIDController := NewGetTipoMantenimientoByIDController(getByIDUseCase)
	updateController := NewUpdateTipoMantenimientoController(updateUseCase)
	deleteController := NewDeleteTipoMantenimientoController(deleteUseCase)
	getAllController := NewGetAllTiposMantenimientoController(getAllUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController
}