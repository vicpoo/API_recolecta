package ag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client llama al servicio algoritmo_genetico_rutas (POST /optimizar).
type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		http: &http.Client{
			Timeout: 90 * time.Second,
		},
	}
}

type Punto struct {
	ID     string  `json:"id"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Nombre string  `json:"nombre,omitempty"`
}

type Base struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Nombre string  `json:"nombre"`
}

type Bloqueo struct {
	ID      string  `json:"id,omitempty"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Reporte string  `json:"reporte,omitempty"`
}

type OptimizarRequest struct {
	Puntos       []Punto   `json:"puntos"`
	BaseInicio   Base      `json:"base_inicio"`
	BaseFin      Base      `json:"base_fin"`
	Bloqueos     []Bloqueo `json:"bloqueos"`
	RadioBloqueo float64   `json:"radio_bloqueo"`
}

type Paso struct {
	Calle       string      `json:"calle"`
	Instruccion string      `json:"instruccion"`
	DistanciaM  float64     `json:"distancia_m"`
	CoordInicio interface{} `json:"coord_inicio"`
}

type Segmento struct {
	De         string `json:"de"`
	A          string `json:"a"`
	DistanciaM float64 `json:"distancia_m"`
	Pasos      []Paso `json:"pasos"`
}

type OptimizarResponse struct {
	BaseInicio         Base        `json:"base_inicio"`
	BaseFin            Base        `json:"base_fin"`
	DistanciaTotalKm   float64     `json:"distancia_total_km"`
	DistanciaTotalM    float64     `json:"distancia_total_m"`
	PuntosInaccesibles []Punto     `json:"puntos_inaccesibles"`
	Segmentos          []Segmento  `json:"segmentos"`
	TodasLasCoords     interface{} `json:"todas_las_coords"`
}

func (c *Client) Optimizar(ctx context.Context, req OptimizarRequest) (*OptimizarResponse, error) {
	if req.RadioBloqueo <= 0 {
		req.RadioBloqueo = 25.0
	}

	buf, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("serializar payload AG: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/optimizar", bytes.NewReader(buf))
	if err != nil {
		return nil, fmt.Errorf("construir request AG: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("conexion con AG: %w", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("AG respondio %d: %s", res.StatusCode, strings.TrimSpace(string(body)))
	}

	var out OptimizarResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("interpretar respuesta AG: %w", err)
	}

	return &out, nil
}

func (c *Client) Health(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/health", nil)
	if err != nil {
		return err
	}

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("AG health %d: %s", res.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}
