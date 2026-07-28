// pipeline_retry_worker.go
package infrastructure

import (
	"log"
	"time"

	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
)

// PipelineRetryWorker es la red de seguridad del pipeline modelo_reportes ->
// clasificador_reportes.
//
// CreateAnomaliaUseCase ya dispara un goroutine que corre el pipeline al
// vuelo apenas se crea el reporte (camino rapido, confirma rapido al
// ciudadano/conductor). El problema es que ese goroutine vive y muere con el
// proceso: si el backend se reinicia a la mitad (Air en dev, un redeploy en
// prod, un crash) el reporte queda huerfano en estado_pipeline = 'pendiente'
// para siempre, sin que nadie lo reintente.
//
// Este worker corre en background durante toda la vida del proceso
// (arrancado una sola vez con `go worker.Run()`, como worker de larga duración) y
// cada `intervalo` revisa la tabla anomalia por filas reclamables:
// pendientes nunca procesadas, 'procesando' abandonadas hace rato (el
// goroutine que las tomo murio antes de terminar), o en 'error' con
// reintentos disponibles. Llama al mismo ProcesarPipelineAnomaliaUseCase.Run
// de siempre -- el claim atomico (IAnomalia.ReclamarPipeline) es lo que
// evita que el camino rapido y este worker procesen la misma anomalia dos
// veces, sin necesidad de locks explicitos ni una cola/broker aparte.
type PipelineRetryWorker struct {
	repo        repositories.IAnomalia
	pipelineUC  *application.ProcesarPipelineAnomaliaUseCase
	intervalo   time.Duration
	maxIntentos int
	loteMaximo  int
}

func NewPipelineRetryWorker(repo repositories.IAnomalia, pipelineUC *application.ProcesarPipelineAnomaliaUseCase) *PipelineRetryWorker {
	return &PipelineRetryWorker{
		repo:        repo,
		pipelineUC:  pipelineUC,
		intervalo:   30 * time.Second,
		maxIntentos: application.MaxIntentosPipeline,
		loteMaximo:  20,
	}
}

// Run bloquea: se espera que se dispare con `go worker.Run()` una sola vez
// al arrancar la app (ver anomalia_routes.go), nunca directamente.
func (w *PipelineRetryWorker) Run() {
	log.Println("pipeline retry worker: iniciado, revisando cada", w.intervalo)

	ticker := time.NewTicker(w.intervalo)
	defer ticker.Stop()

	for range ticker.C {
		w.tick()
	}
}

func (w *PipelineRetryWorker) tick() {
	anomalias, err := w.repo.ListoParaPipeline(w.maxIntentos, application.PipelineProcesandoStaleDespues, w.loteMaximo)
	if err != nil {
		log.Println("pipeline retry worker: error al listar anomalias pendientes:", err)
		return
	}
	if len(anomalias) == 0 {
		return
	}

	log.Println("pipeline retry worker: reintentando", len(anomalias), "anomalia(s)")

	for _, a := range anomalias {
		origen := "ciudadano"
		if a.ConductorID != nil {
			origen = "conductor"
		}

		// Secuencial a proposito: el volumen esperado (reportes de calle) es
		// bajo, y procesar uno por uno evita mandar rafagas concurrentes a
		// modelo_reportes/clasificador_reportes en cada tick. Run() vuelve a
		// intentar el claim atomico, asi que si el camino rapido ya proceso
		// esta misma anomalia entre el listado y este punto, simplemente no
		// hace nada.
		w.pipelineUC.Run(a.AnomaliaID, a.Descripcion, 1, &origen)
	}
}
