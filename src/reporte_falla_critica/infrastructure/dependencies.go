// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
)

func InitReporteFallaCriticaDependencies() (
	*CreateReporteFallaCriticaController,
	*GetReporteFallaCriticaByIdController,
	*UpdateReporteFallaCriticaController,
	*DeleteReporteFallaCriticaController,
	*GetAllReportesFallaCriticaController,
	*GetReportesFallaCriticaByCamionIDController,
	*GetReportesFallaCriticaByConductorIDController,
	*GetReportesFallaCriticaByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresReporteFallaCriticaRepository()

	// Casos de uso
	createUseCase := application.NewCreateReporteFallaCriticaUseCase(repo)
	getByIDUseCase := application.NewGetReporteFallaCriticaByIdUseCase(repo)
	updateUseCase := application.NewUpdateReporteFallaCriticaUseCase(repo)
	deleteUseCase := application.NewDeleteReporteFallaCriticaUseCase(repo)
	getAllUseCase := application.NewGetAllReportesFallaCriticaUseCase(repo)
	getByCamionIDUseCase := application.NewGetReportesFallaCriticaByCamionIDUseCase(repo)
	getByConductorIDUseCase := application.NewGetReportesFallaCriticaByConductorIDUseCase(repo)
	getByFechaRangeUseCase := application.NewGetReportesFallaCriticaByFechaRangeUseCase(repo)

	// Controladores
	createController := NewCreateReporteFallaCriticaController(createUseCase)
	getByIDController := NewGetReporteFallaCriticaByIdController(getByIDUseCase)
	updateController := NewUpdateReporteFallaCriticaController(updateUseCase)
	deleteController := NewDeleteReporteFallaCriticaController(deleteUseCase)
	getAllController := NewGetAllReportesFallaCriticaController(getAllUseCase)
	getByCamionIDController := NewGetReportesFallaCriticaByCamionIDController(getByCamionIDUseCase)
	getByConductorIDController := NewGetReportesFallaCriticaByConductorIDController(getByConductorIDUseCase)
	getByFechaRangeController := NewGetReportesFallaCriticaByFechaRangeController(getByFechaRangeUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController, getByCamionIDController, getByConductorIDController, getByFechaRangeController
}