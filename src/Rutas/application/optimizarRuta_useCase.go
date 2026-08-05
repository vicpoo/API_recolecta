package application

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
	agclient "github.com/vicpoo/API_recolecta/src/Rutas/infrastructure/ag"
)

var puntoIDFromLabel = regexp.MustCompile(`^(\d+)\s*-`)

type OptimizarRutaUseCase struct {
	rutaRepo ports.IRuta
	puntoRepo ports.IPuntoRecoleccion
	ag       *agclient.Client
}

func NewOptimizarRutaUseCase(
	rutaRepo ports.IRuta,
	puntoRepo ports.IPuntoRecoleccion,
	ag *agclient.Client,
) *OptimizarRutaUseCase {
	return &OptimizarRutaUseCase{
		rutaRepo:  rutaRepo,
		puntoRepo: puntoRepo,
		ag:        ag,
	}
}

type OptimizarRutaInput struct {
	Bloqueos     []agclient.Bloqueo
	RadioBloqueo float64
}

type OptimizarRutaResult struct {
	RutaID           int32                  `json:"ruta_id"`
	Nombre           string                 `json:"nombre"`
	DistanciaTotalKm float64                `json:"distancia_total_km"`
	DistanciaTotalM  float64                `json:"distancia_total_m"`
	Puntos           []entities.PuntoRecoleccion `json:"puntos"`
	Segmentos        []agclient.Segmento    `json:"segmentos"`
	Coordenadas      interface{}            `json:"coordenadas"`
}

func (uc *OptimizarRutaUseCase) Run(
	ctx context.Context,
	tenantID int,
	rutaID int32,
	input OptimizarRutaInput,
) (*OptimizarRutaResult, error) {
	ruta, err := uc.rutaRepo.GetById(ctx, tenantID, rutaID)
	if err != nil {
		return nil, fmt.Errorf("ruta no encontrada: %w", err)
	}

	puntos, err := uc.puntoRepo.GetByRuta(ctx, tenantID, rutaID)
	if err != nil {
		return nil, fmt.Errorf("obtener puntos: %w", err)
	}
	if len(puntos) < 2 {
		return nil, fmt.Errorf("la ruta necesita al menos base inicio y base fin")
	}

	baseInicio, baseFin := basesFromJSON(ruta.JsonRuta, puntos)
	puntoInicio := puntos[0]
	puntoFin := puntos[len(puntos)-1]

	intermedios := puntos[1 : len(puntos)-1]
	if len(intermedios) == 0 {
		return nil, fmt.Errorf("se necesita al menos un punto intermedio ademas de las bases")
	}

	puntosAG := make([]agclient.Punto, 0, len(intermedios))
	for _, p := range intermedios {
		if p.Lat == 0 && p.Lon == 0 {
			return nil, fmt.Errorf("punto %d sin coordenadas en Redis", p.PuntoID)
		}
		nombre := p.CP
		if nombre == "" {
			nombre = fmt.Sprintf("Punto %d", p.PuntoID)
		}
		puntosAG = append(puntosAG, agclient.Punto{
			ID:     strconv.Itoa(int(p.PuntoID)),
			Lat:    p.Lat,
			Lng:    p.Lon,
			Nombre: nombre,
		})
	}

	bloqueos := input.Bloqueos
	if bloqueos == nil {
		bloqueos = []agclient.Bloqueo{}
	}

	agResp, err := uc.ag.Optimizar(ctx, agclient.OptimizarRequest{
		Puntos:       puntosAG,
		BaseInicio:   baseInicio,
		BaseFin:      baseFin,
		Bloqueos:     bloqueos,
		RadioBloqueo: input.RadioBloqueo,
	})
	if err != nil {
		return nil, fmt.Errorf("optimizacion AG: %w", err)
	}

	jsonRuta, err := mergeJsonRutaOptimizada(ruta.JsonRuta, agResp, puntos, puntoInicio, puntoFin)
	if err != nil {
		return nil, err
	}

	ruta.JsonRuta = jsonRuta
	if err := uc.rutaRepo.Update(ctx, tenantID, ruta); err != nil {
		return nil, fmt.Errorf("guardar ruta optimizada: %w", err)
	}

	puntosActualizados, err := uc.puntoRepo.GetByRuta(ctx, tenantID, rutaID)
	if err != nil {
		return nil, fmt.Errorf("obtener puntos actualizados: %w", err)
	}

	return &OptimizarRutaResult{
		RutaID:           ruta.RutaID,
		Nombre:           ruta.Nombre,
		DistanciaTotalKm: agResp.DistanciaTotalKm,
		DistanciaTotalM:  agResp.DistanciaTotalM,
		Puntos:           puntosActualizados,
		Segmentos:        agResp.Segmentos,
		Coordenadas:      agResp.TodasLasCoords,
	}, nil
}

func basesFromJSON(jsonRuta string, puntos []entities.PuntoRecoleccion) (agclient.Base, agclient.Base) {
	inicio := agclient.Base{
		Lat:    puntos[0].Lat,
		Lng:    puntos[0].Lon,
		Nombre: labelForPunto(puntos[0], "Base Inicio"),
	}
	fin := agclient.Base{
		Lat:    puntos[len(puntos)-1].Lat,
		Lng:    puntos[len(puntos)-1].Lon,
		Nombre: labelForPunto(puntos[len(puntos)-1], "Base Fin"),
	}

	if jsonRuta == "" {
		return inicio, fin
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonRuta), &parsed); err != nil {
		return inicio, fin
	}

	if bi, ok := parsed["base_inicio"].(map[string]interface{}); ok {
		inicio = baseFromMap(bi, inicio.Nombre)
	}
	if bf, ok := parsed["base_fin"].(map[string]interface{}); ok {
		fin = baseFromMap(bf, fin.Nombre)
	}

	return inicio, fin
}

func baseFromMap(m map[string]interface{}, fallbackNombre string) agclient.Base {
	b := agclient.Base{Nombre: fallbackNombre}
	if v, ok := m["nombre"].(string); ok && v != "" {
		b.Nombre = v
	}
	if v, ok := toFloat(m["lat"]); ok {
		b.Lat = v
	}
	if v, ok := toFloat(m["lng"]); ok {
		b.Lng = v
	}
	if v, ok := toFloat(m["lon"]); ok && b.Lng == 0 {
		b.Lng = v
	}
	return b
}

func labelForPunto(p entities.PuntoRecoleccion, fallback string) string {
	if p.CP != "" {
		return p.CP
	}
	return fallback
}

func mergeJsonRutaOptimizada(
	existingJSON string,
	agResp *agclient.OptimizarResponse,
	puntos []entities.PuntoRecoleccion,
	puntoInicio, puntoFin entities.PuntoRecoleccion,
) (string, error) {
	meta := map[string]interface{}{
		"type":        "LineString",
		"coordinates": coordsToGeoJSON(agResp.TodasLasCoords),
		"puntos":      []interface{}{},
	}

	if existingJSON != "" {
		var parsed map[string]interface{}
		if err := json.Unmarshal([]byte(existingJSON), &parsed); err == nil {
			for _, k := range []string{"zona", "turno", "descripcion"} {
				if v, ok := parsed[k]; ok {
					meta[k] = v
				}
			}
			if v, ok := parsed["puntos"]; ok {
				meta["puntos"] = v
			}
		}
	}

	meta["base_inicio"] = agResp.BaseInicio
	meta["base_fin"] = agResp.BaseFin
	meta["optimizada"] = true
	meta["distancia_total_km"] = agResp.DistanciaTotalKm
	meta["ruta_optimizada_coords"] = agResp.TodasLasCoords

	if orden := ordenFromSegmentos(agResp.Segmentos); len(orden) > 0 {
		if reordered := reorderPuntosJSON(meta["puntos"], puntos, orden, puntoInicio, puntoFin); reordered != nil {
			meta["puntos"] = reordered
		}
	}

	raw, err := json.Marshal(meta)
	if err != nil {
		return "", fmt.Errorf("serializar json_ruta optimizada: %w", err)
	}
	return string(raw), nil
}

func coordsToGeoJSON(raw interface{}) [][]float64 {
	items, ok := raw.([]interface{})
	if !ok {
		return nil
	}

	out := make([][]float64, 0, len(items))
	for _, item := range items {
		switch c := item.(type) {
		case []interface{}:
			if len(c) >= 2 {
				lat, _ := toFloat(c[0])
				lng, _ := toFloat(c[1])
				out = append(out, []float64{lng, lat})
			}
		case map[string]interface{}:
			lat, latOK := toFloat(c["lat"])
			lng, lngOK := toFloat(c["lng"])
			if !lngOK {
				lng, lngOK = toFloat(c["lon"])
			}
			if latOK && lngOK {
				out = append(out, []float64{lng, lat})
			}
		}
	}
	return out
}

func ordenFromSegmentos(segmentos []agclient.Segmento) []int32 {
	if len(segmentos) == 0 {
		return nil
	}

	var ids []int32
	if id := idFromLabel(segmentos[0].De); id > 0 {
		ids = append(ids, id)
	}
	for _, seg := range segmentos {
		if id := idFromLabel(seg.A); id > 0 {
			ids = append(ids, id)
		}
	}
	return ids
}

func idFromLabel(label string) int32 {
	m := puntoIDFromLabel.FindStringSubmatch(label)
	if len(m) < 2 {
		return 0
	}
	n, err := strconv.Atoi(m[1])
	if err != nil {
		return 0
	}
	return int32(n)
}

func reorderPuntosJSON(
	existing interface{},
	puntos []entities.PuntoRecoleccion,
	orden []int32,
	puntoInicio, puntoFin entities.PuntoRecoleccion,
) []map[string]interface{} {
	byID := map[int32]map[string]interface{}{}
	if arr, ok := existing.([]interface{}); ok {
		for i, item := range arr {
			m, ok := item.(map[string]interface{})
			if !ok {
				continue
			}
			id := int32(i + 1)
			if v, ok := m["id"]; ok {
				if n, ok := toFloat(v); ok {
					id = int32(n)
				}
			}
			byID[id] = m
		}
	}

	for _, p := range puntos {
		if _, exists := byID[p.PuntoID]; exists {
			continue
		}
		byID[p.PuntoID] = map[string]interface{}{
			"id":     p.PuntoID,
			"nombre": labelForPunto(p, fmt.Sprintf("Punto %d", p.PuntoID)),
			"lat":    p.Lat,
			"lng":    p.Lon,
		}
	}

	out := make([]map[string]interface{}, 0, len(puntos))
	appendPunto := func(id int32, orden int) {
		if m, ok := byID[id]; ok {
			copy := map[string]interface{}{}
			for k, v := range m {
				copy[k] = v
			}
			copy["orden"] = orden
			out = append(out, copy)
		}
	}

	ord := 1
	appendPunto(puntoInicio.PuntoID, ord)
	ord++

	for _, id := range orden {
		if id == puntoInicio.PuntoID || id == puntoFin.PuntoID {
			continue
		}
		appendPunto(id, ord)
		ord++
	}

	appendPunto(puntoFin.PuntoID, ord)
	return out
}

func toFloat(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case json.Number:
		f, err := n.Float64()
		return f, err == nil
	default:
		return 0, false
	}
}
