// anomalia_creada_notifier.go
package infrastructure

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// HTTPAnomaliaCreadaNotifier implementa domain.AnomaliaCreadaNotifier
// mandando un POST a un webhook externo (por ahora, el algoritmo genetico
// de rutas) cada vez que se crea una anomalia con lat/lon.
//
// A diferencia del pipeline modelo_reportes/clasificador_reportes -- que
// usa outbox + retry worker (ver pipeline_retry_worker.go) porque su
// resultado se persiste en la propia fila y el sistema depende de el --
// esto es una notificacion de salida hacia OTRO sistema: si falla, no deja
// nada inconsistente del lado de Recolecta. Por eso, por ahora, es "mejor
// esfuerzo": un solo intento con timeout corto, se loguea si falla, y no se
// reintenta. Si el webhook se vuelve critico (el AG no puede perder
// eventos) esto se puede promover al mismo patron outbox+retry sin tener
// que rediseñar nada -- el puerto (domain.AnomaliaCreadaNotifier) no
// cambiaria.
type HTTPAnomaliaCreadaNotifier struct {
	webhookURL string
	http       *http.Client
}

func NewHTTPAnomaliaCreadaNotifier(webhookURL string) *HTTPAnomaliaCreadaNotifier {
	return &HTTPAnomaliaCreadaNotifier{
		webhookURL: webhookURL,
		http: &http.Client{
			// Timeout corto: esto corre en un goroutine de "fire and forget"
			// disparado desde CreateAnomaliaUseCase, no debe quedar colgado si
			// el webhook (que hoy ni siquiera esta desplegado) no responde.
			Timeout: 5 * time.Second,
		},
	}
}

// anomaliaCreadaPayload es el contrato que pidio el equipo del webhook.
type anomaliaCreadaPayload struct {
	IDAnomalia  int32   `json:"id_anomalia"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Descripcion string  `json:"descripcion"`
	// Status: fijo en "aprobado" por ahora. Recolecta todavia no tiene un
	// flujo de aprobacion propio para anomalias (el campo real es `estado`,
	// con PENDIENTE/EN_PROCESO/RESUELTA -- ver anomalia.go); si mas adelante
	// se agrega un paso de aprobacion, este valor debe salir de ahi en vez
	// de estar fijo.
	Status string `json:"status"`
}

func (n *HTTPAnomaliaCreadaNotifier) Notificar(anomaliaID int32, lat, lon float64, descripcion string) {
	payload := anomaliaCreadaPayload{
		IDAnomalia:  anomaliaID,
		Lat:         lat,
		Lng:         lon,
		Descripcion: descripcion,
		Status:      "aprobado",
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		log.Println("anomalia_creada webhook: no se pudo serializar el payload de la anomalia", anomaliaID, ":", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, n.webhookURL, bytes.NewReader(buf))
	if err != nil {
		log.Println("anomalia_creada webhook: no se pudo construir el request para la anomalia", anomaliaID, ":", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := n.http.Do(req)
	if err != nil {
		log.Println("anomalia_creada webhook: fallo el envio para la anomalia", anomaliaID, "(", n.webhookURL, "):", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Println("anomalia_creada webhook: respuesta", res.StatusCode, "al notificar la anomalia", anomaliaID)
		return
	}

	log.Println("anomalia_creada webhook: anomalia", anomaliaID, "notificada correctamente")
}
