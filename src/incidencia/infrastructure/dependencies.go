// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
)

func InitIncidenciaDependencies() (
	*CreateIncidenciaController,
	*GetIncidenciaByIDController,
	*UpdateIncidenciaController,
	*DeleteIncidenciaController,
	*GetAllIncidenciasController,
	*GetIncidenciasByConductorIDController,
	*GetIncidenciasByPuntoRecoleccionIDController,
	*GetIncidenciasByFechaRangeController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresIncidenciaRepository()

	// Casos de uso CRUD básicos
	createUseCase := application.NewCreateIncidenciaUseCase(repo)
	getByIDUseCase := application.NewGetIncidenciaByIDUseCase(repo)
	updateUseCase := application.NewUpdateIncidenciaUseCase(repo)
	deleteUseCase := application.NewDeleteIncidenciaUseCase(repo)
	getAllUseCase := application.NewGetAllIncidenciasUseCase(repo)
	
	// Casos de uso específicos
	getByConductorIDUseCase := application.NewGetIncidenciasByConductorIDUseCase(repo)
	getByPuntoRecoleccionIDUseCase := application.NewGetIncidenciasByPuntoRecoleccionIDUseCase(repo)
	getByFechaRangeUseCase := application.NewGetIncidenciasByFechaRangeUseCase(repo)

	// Controladores CRUD básicos
	createController := NewCreateIncidenciaController(createUseCase)
	getByIDController := NewGetIncidenciaByIDController(getByIDUseCase)
	updateController := NewUpdateIncidenciaController(updateUseCase)
	deleteController := NewDeleteIncidenciaController(deleteUseCase)
	getAllController := NewGetAllIncidenciasController(getAllUseCase)
	
	// Controladores específicos
	getByConductorIDController := NewGetIncidenciasByConductorIDController(getByConductorIDUseCase)
	getByPuntoRecoleccionIDController := NewGetIncidenciasByPuntoRecoleccionIDController(getByPuntoRecoleccionIDUseCase)
	getByFechaRangeController := NewGetIncidenciasByFechaRangeController(getByFechaRangeUseCase)

	return createController, 
	       getByIDController, 
	       updateController, 
	       deleteController, 
	       getAllController,
	       getByConductorIDController,
	       getByPuntoRecoleccionIDController,
	       getByFechaRangeController
}