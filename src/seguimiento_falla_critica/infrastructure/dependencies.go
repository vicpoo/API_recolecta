// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/application"
)

func InitSeguimientoFallaCriticaDependencies() (
	*CreateSeguimientoFallaCriticaController,
	*GetSeguimientoFallaCriticaByIdController,
	*UpdateSeguimientoFallaCriticaController,
	*DeleteSeguimientoFallaCriticaController,
	*GetAllSeguimientosFallaCriticaController,
	*GetSeguimientosFallaCriticaByFallaIDController,
	*GetSeguimientosFallaCriticaByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresSeguimientoFallaCriticaRepository()

	// Casos de uso
	createUseCase := application.NewCreateSeguimientoFallaCriticaUseCase(repo)
	getByIDUseCase := application.NewGetSeguimientoFallaCriticaByIdUseCase(repo)
	updateUseCase := application.NewUpdateSeguimientoFallaCriticaUseCase(repo)
	deleteUseCase := application.NewDeleteSeguimientoFallaCriticaUseCase(repo)
	getAllUseCase := application.NewGetAllSeguimientosFallaCriticaUseCase(repo)
	getByFallaIDUseCase := application.NewGetSeguimientosFallaCriticaByFallaIDUseCase(repo)
	getByFechaRangeUseCase := application.NewGetSeguimientosFallaCriticaByFechaRangeUseCase(repo)

	// Controladores
	createController := NewCreateSeguimientoFallaCriticaController(createUseCase)
	getByIDController := NewGetSeguimientoFallaCriticaByIdController(getByIDUseCase)
	updateController := NewUpdateSeguimientoFallaCriticaController(updateUseCase)
	deleteController := NewDeleteSeguimientoFallaCriticaController(deleteUseCase)
	getAllController := NewGetAllSeguimientosFallaCriticaController(getAllUseCase)
	getByFallaIDController := NewGetSeguimientosFallaCriticaByFallaIDController(getByFallaIDUseCase)
	getByFechaRangeController := NewGetSeguimientosFallaCriticaByFechaRangeController(getByFechaRangeUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController, getByFallaIDController, getByFechaRangeController
}