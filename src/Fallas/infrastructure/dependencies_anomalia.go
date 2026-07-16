// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

func InitAnomaliaDependencies() (
	*CreateAnomaliaController,
	*GetAnomaliaByIdController,
	*UpdateAnomaliaController,
	*DeleteAnomaliaController,
	*GetAllAnomaliasController,
	*GetAnomaliasByPuntoIDController,
	*GetAnomaliasByChoferIDController,
	*GetAnomaliasByCamionIDController,
	*GetAnomaliasByRutaIDController,
	*GetAnomaliasByReferenciaIDController,
	*GetAnomaliasByEstadoController,
	*GetAnomaliasByTipoAnomaliaController,
	*GetAnomaliasByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresAnomaliaRepository()

	// Casos de uso
	createUseCase := application.NewCreateAnomaliaUseCase(repo)
	getByIDUseCase := application.NewGetAnomaliaByIdUseCase(repo)
	updateUseCase := application.NewUpdateAnomaliaUseCase(repo)
	deleteUseCase := application.NewDeleteAnomaliaUseCase(repo)
	getAllUseCase := application.NewGetAllAnomaliasUseCase(repo)
	getByPuntoIDUseCase := application.NewGetAnomaliasByPuntoIDUseCase(repo)
	getByChoferIDUseCase := application.NewGetAnomaliasByConductorIDUseCase(repo)
	getByCamionIDUseCase := application.NewGetAnomaliasByCamionIDUseCase(repo)
	getByRutaIDUseCase := application.NewGetAnomaliasByRutaIDUseCase(repo)
	getByReferenciaIDUseCase := application.NewGetAnomaliasByReferenciaIDUseCase(repo)
	getByEstadoUseCase := application.NewGetAnomaliasByEstadoUseCase(repo)
	getByTipoAnomaliaUseCase := application.NewGetAnomaliasByTipoAnomaliaUseCase(repo)
	getByFechaRangeUseCase := application.NewGetAnomaliasByFechaRangeUseCase(repo)

	// Controladores
	createController := NewCreateAnomaliaController(createUseCase)
	getByIDController := NewGetAnomaliaByIdController(getByIDUseCase)
	updateController := NewUpdateAnomaliaController(updateUseCase)
	deleteController := NewDeleteAnomaliaController(deleteUseCase)
	getAllController := NewGetAllAnomaliasController(getAllUseCase)
	getByPuntoIDController := NewGetAnomaliasByPuntoIDController(getByPuntoIDUseCase)
	getByChoferIDController := NewGetAnomaliasByChoferIDController(getByChoferIDUseCase)
	getByCamionIDController := NewGetAnomaliasByCamionIDController(getByCamionIDUseCase)
	getByRutaIDController := NewGetAnomaliasByRutaIDController(getByRutaIDUseCase)
	getByReferenciaIDController := NewGetAnomaliasByReferenciaIDController(getByReferenciaIDUseCase)
	getByEstadoController := NewGetAnomaliasByEstadoController(getByEstadoUseCase)
	getByTipoAnomaliaController := NewGetAnomaliasByTipoAnomaliaController(getByTipoAnomaliaUseCase)
	getByFechaRangeController := NewGetAnomaliasByFechaRangeController(getByFechaRangeUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController,
		getByPuntoIDController, getByChoferIDController, getByCamionIDController, getByRutaIDController,
		getByReferenciaIDController, getByEstadoController, getByTipoAnomaliaController, getByFechaRangeController
}