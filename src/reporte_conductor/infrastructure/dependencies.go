// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
)

func InitReporteConductorDependencies() (
	*CreateReporteConductorController,
	*GetReporteConductorByIdController,
	*UpdateReporteConductorController,
	*DeleteReporteConductorController,
	*GetAllReportesConductorController,
	*GetReportesConductorByCamionIDController,
	*GetReportesConductorByConductorIDController,
	*GetReportesConductorByRutaIDController,
	*GetReportesConductorByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresReporteConductorRepository()

	// Casos de uso
	createUseCase := application.NewCreateReporteConductorUseCase(repo)
	getByIDUseCase := application.NewGetReporteConductorByIdUseCase(repo)
	updateUseCase := application.NewUpdateReporteConductorUseCase(repo)
	deleteUseCase := application.NewDeleteReporteConductorUseCase(repo)
	getAllUseCase := application.NewGetAllReportesConductorUseCase(repo)
	getByCamionIDUseCase := application.NewGetReportesConductorByCamionIDUseCase(repo)
	getByConductorIDUseCase := application.NewGetReportesConductorByConductorIDUseCase(repo)
	getByRutaIDUseCase := application.NewGetReportesConductorByRutaIDUseCase(repo)
	getByFechaRangeUseCase := application.NewGetReportesConductorByFechaRangeUseCase(repo)

	// Controladores
	createController := NewCreateReporteConductorController(createUseCase)
	getByIDController := NewGetReporteConductorByIdController(getByIDUseCase)
	updateController := NewUpdateReporteConductorController(updateUseCase)
	deleteController := NewDeleteReporteConductorController(deleteUseCase)
	getAllController := NewGetAllReportesConductorController(getAllUseCase)
	getByCamionIDController := NewGetReportesConductorByCamionIDController(getByCamionIDUseCase)
	getByConductorIDController := NewGetReportesConductorByConductorIDController(getByConductorIDUseCase)
	getByRutaIDController := NewGetReportesConductorByRutaIDController(getByRutaIDUseCase)
	getByFechaRangeController := NewGetReportesConductorByFechaRangeController(getByFechaRangeUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController, getByCamionIDController, getByConductorIDController, getByRutaIDController, getByFechaRangeController
}