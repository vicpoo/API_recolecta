// anomalia_repository.go
package domain

import (
	"time"

	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
)

type IAnomalia interface {
	// Operaciones CRUD básicas
	Save(anomalia *entities.Anomalia) error
	Update(anomalia *entities.Anomalia) error
	Delete(id int32) error
	GetAll() ([]entities.Anomalia, error)
	GetByID(id int32) (*entities.Anomalia, error)

	// Métodos específicos para Anomalia (cubren los antiguos dominios
	// Incidencia, ReporteConductor, ReporteFallaCritica y
	// SeguimientoFallaCritica, ahora unificados por TipoAnomalia)
	GetByTipoAnomalia(tipoAnomalia entities.TipoAnomalia) ([]entities.Anomalia, error)
	GetByEstado(estado string) ([]entities.Anomalia, error)
	GetByPuntoID(puntoID int32) ([]entities.Anomalia, error)
	GetByConductorID(conductorID int32) ([]entities.Anomalia, error)
	GetByCiudadanoID(ciudadanoID int32) ([]entities.Anomalia, error)
	GetByCamionID(camionID int32) ([]entities.Anomalia, error)
	GetByRutaID(rutaID int32) ([]entities.Anomalia, error)
	GetByReferenciaID(referenciaID int32) ([]entities.Anomalia, error)
	GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Anomalia, error)

	// ActualizarPipeline persiste el resultado del pipeline
	// modelo_reportes -> clasificador_reportes sobre una anomalia ya
	// existente. Es una actualizacion dirigida (no pasa por Update/la
	// entidad completa) para no interferir con el flujo CRUD normal.
	// Los punteros nil dejan la columna correspondiente sin tocar.
	ActualizarPipeline(anomaliaID int32, estadoPipeline string, nivelRiesgo *string, inferenciaID *int32, categoria *string, subtipo *string, accion *string, pipelineError *string) error

	// ReclamarPipeline es un "claim" atomico: intenta tomar la fila para
	// procesarla (pasa a estado_pipeline = 'procesando' e incrementa
	// pipeline_intentos) solo si esta en un estado reclamable:
	//   - 'pendiente' (recien creada, nunca se proceso)
	//   - 'procesando' pero abandonada hace mas de procesandoStaleDespues
	//     (el proceso que la tomo se cayo/reinicio a la mitad)
	//   - 'error' con pipeline_intentos < maxIntentos (reintento acotado)
	// Devuelve false (sin error) si la fila no era reclamable -- por
	// ejemplo porque otro disparo (el goroutine del alta o un tick del
	// worker) ya la tomo. Este claim es lo que evita procesar la misma
	// anomalia dos veces cuando el camino rapido y el worker de reintento
	// coinciden.
	ReclamarPipeline(anomaliaID int32, maxIntentos int, procesandoStaleDespues time.Duration) (bool, error)

	// ListoParaPipeline devuelve anomalias en un estado reclamable (ver
	// ReclamarPipeline) para que PipelineRetryWorker las vuelva a intentar.
	ListoParaPipeline(maxIntentos int, procesandoStaleDespues time.Duration, limit int) ([]entities.Anomalia, error)
}
