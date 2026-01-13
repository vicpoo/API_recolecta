// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

func InitAlertaMantenimientoDependencies() (
	*CreateAlertaMantenimientoController,
	*GetAlertaMantenimientoByIDController,
	*UpdateAlertaMantenimientoController,
	*DeleteAlertaMantenimientoController,
	*GetAllAlertasMantenimientoController,
	*GetAlertasByCamionIDController,
	*GetAlertasByTipoMantenimientoIDController,
	*GetAlertasPendientesController,
	*GetAlertasAtendidasController,
	*MarcarAlertaAtendidaController,
	*GetAlertasByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresAlertaMantenimientoRepository()

	// Casos de uso
	createUseCase := application.NewCreateAlertaMantenimientoUseCase(repo)
	getByIDUseCase := application.NewGetAlertaMantenimientoByIDUseCase(repo)
	updateUseCase := application.NewUpdateAlertaMantenimientoUseCase(repo)
	deleteUseCase := application.NewDeleteAlertaMantenimientoUseCase(repo)
	getAllUseCase := application.NewGetAllAlertasMantenimientoUseCase(repo)
	getByCamionUseCase := application.NewGetAlertasByCamionIDUseCase(repo)
	getByTipoUseCase := application.NewGetAlertasByTipoMantenimientoIDUseCase(repo)
	getPendientesUseCase := application.NewGetAlertasPendientesUseCase(repo)
	getAtendidasUseCase := application.NewGetAlertasAtendidasUseCase(repo)
	marcarAtendidaUseCase := application.NewMarcarAlertaAtendidaUseCase(repo)
	getByFechaRangeUseCase := application.NewGetAlertasByFechaRangeUseCase(repo)

	// Controladores
	createController := NewCreateAlertaMantenimientoController(createUseCase)
	getByIDController := NewGetAlertaMantenimientoByIDController(getByIDUseCase)
	updateController := NewUpdateAlertaMantenimientoController(updateUseCase)
	deleteController := NewDeleteAlertaMantenimientoController(deleteUseCase)
	getAllController := NewGetAllAlertasMantenimientoController(getAllUseCase)
	getByCamionController := NewGetAlertasByCamionIDController(getByCamionUseCase)
	getByTipoController := NewGetAlertasByTipoMantenimientoIDController(getByTipoUseCase)
	getPendientesController := NewGetAlertasPendientesController(getPendientesUseCase)
	getAtendidasController := NewGetAlertasAtendidasController(getAtendidasUseCase)
	marcarAtendidaController := NewMarcarAlertaAtendidaController(marcarAtendidaUseCase)
	getByFechaRangeController := NewGetAlertasByFechaRangeController(getByFechaRangeUseCase)

	return createController, getByIDController, updateController, deleteController, getAllController, getByCamionController, getByTipoController, getPendientesController, getAtendidasController, marcarAtendidaController, getByFechaRangeController
}