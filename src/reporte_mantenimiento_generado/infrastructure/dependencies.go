// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/application"
)

func InitReporteMantenimientoGeneradoDependencies() (
	*CreateReporteMantenimientoGeneradoController,
	*GetReporteMantenimientoGeneradoByIdController,
	*UpdateReporteMantenimientoGeneradoController,
	*DeleteReporteMantenimientoGeneradoController,
	*GetAllReportesMantenimientoGeneradoController,
	*GetReportesMantenimientoGeneradoByCoordinadorIDController,
	*GetReportesMantenimientoGeneradoByFechaRangeController,
	*GetReportesMantenimientoGeneradoByFechaGeneracionRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresReporteMantenimientoGeneradoRepository()

	// Casos de uso
	createUseCase := application.NewCreateReporteMantenimientoGeneradoUseCase(repo)
	getByIDUseCase := application.NewGetReporteMantenimientoGeneradoByIdUseCase(repo)
	updateUseCase := application.NewUpdateReporteMantenimientoGeneradoUseCase(repo)
	deleteUseCase := application.NewDeleteReporteMantenimientoGeneradoUseCase(repo)
	getAllUseCase := application.NewGetAllReportesMantenimientoGeneradoUseCase(repo)
	getByCoordinadorIDUseCase := application.NewGetReportesMantenimientoGeneradoByCoordinadorIDUseCase(repo)
	getByFechaRangeUseCase := application.NewGetReportesMantenimientoGeneradoByFechaRangeUseCase(repo)
	getByFechaGeneracionRangeUseCase := application.NewGetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase(repo)

	// Controladores
	createController := NewCreateReporteMantenimientoGeneradoController(createUseCase)
	getByIDController := NewGetReporteMantenimientoGeneradoByIdController(getByIDUseCase)
	updateController := NewUpdateReporteMantenimientoGeneradoController(updateUseCase)
	deleteController := NewDeleteReporteMantenimientoGeneradoController(deleteUseCase)
	getAllController := NewGetAllReportesMantenimientoGeneradoController(getAllUseCase)
	getByCoordinadorIDController := NewGetReportesMantenimientoGeneradoByCoordinadorIDController(getByCoordinadorIDUseCase)
	getByFechaRangeController := NewGetReportesMantenimientoGeneradoByFechaRangeController(getByFechaRangeUseCase)
	getByFechaGeneracionRangeController := NewGetReportesMantenimientoGeneradoByFechaGeneracionRangeController(getByFechaGeneracionRangeUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController, getByCoordinadorIDController, getByFechaRangeController, getByFechaGeneracionRangeController
}