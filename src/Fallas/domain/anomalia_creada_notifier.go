// anomalia_creada_notifier.go
package domain

// AnomaliaCreadaNotifier es el puerto (arquitectura hexagonal, igual que
// PipelineReportesClient en pipeline_reportes.go) para avisar a un sistema
// externo cuando se crea una anomalia con coordenadas conocidas.
//
// Caso concreto actual: el equipo del algoritmo genetico de rutas necesita
// enterarse de cada anomalia (bache, calle bloqueada, etc.) para recalcular
// el grafo de rutas. Ver HTTPAnomaliaCreadaNotifier en infrastructure para
// la implementacion real (POST a un webhook).
type AnomaliaCreadaNotifier interface {
	// Notificar avisa de una anomalia recien creada. No devuelve error: se
	// piensa para llamarse en background (goroutine) sin bloquear la
	// respuesta al ciudadano/conductor -- ver CreateAnomaliaUseCase.Run.
	Notificar(anomaliaID int32, lat, lon float64, descripcion string)
}
