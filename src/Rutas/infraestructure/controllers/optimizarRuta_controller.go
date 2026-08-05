package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
	agclient "github.com/vicpoo/API_recolecta/src/Rutas/infrastructure/ag"
	"github.com/vicpoo/API_recolecta/src/core"
)

type OptimizarRutaController struct {
	uc *application.OptimizarRutaUseCase
}

func NewOptimizarRutaController(uc *application.OptimizarRutaUseCase) *OptimizarRutaController {
	return &OptimizarRutaController{uc: uc}
}

// @Summary      Optimizar ruta con algoritmo genetico
// @Tags         Ruta
// @Accept       json
// @Produce      json
// @Param        id path int true "ID de la ruta"
// @Param        body body object false "Bloqueos opcionales"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} core.ErrorResponse
// @Failure      404 {object} core.ErrorResponse
// @Failure      500 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/rutas/{id}/optimizar [post]
func (c *OptimizarRutaController) Run(ctx *gin.Context) {
	tenantID, ok := core.TenantIDFromContext(ctx)
	if !ok {
		core.RespondBadRequest(ctx, "tenant no encontrado en token", nil)
		return
	}

	rutaID, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil || rutaID <= 0 {
		core.RespondBadRequest(ctx, "id de ruta invalido", nil)
		return
	}

	var body struct {
		Bloqueos     []agclient.Bloqueo `json:"bloqueos"`
		RadioBloqueo float64            `json:"radio_bloqueo"`
	}
	_ = ctx.ShouldBindJSON(&body)

	result, err := c.uc.Run(ctx.Request.Context(), tenantID, int32(rutaID), application.OptimizarRutaInput{
		Bloqueos:     body.Bloqueos,
		RadioBloqueo: body.RadioBloqueo,
	})
	if err != nil {
		msg := err.Error()
		switch {
		case strings.Contains(msg, "ruta no encontrada"):
			core.RespondNotFound(ctx, "Ruta", ctx.Param("id"))
		case strings.Contains(msg, "necesita") || strings.Contains(msg, "intermedio") || strings.Contains(msg, "coordenadas"):
			core.RespondBadRequest(ctx, msg, nil)
		default:
			core.RespondInternalServerError(ctx, "No se pudo optimizar la ruta", err)
		}
		return
	}

	core.RespondOK(ctx, gin.H{
		"success": true,
		"message": "Ruta optimizada exitosamente",
		"data":    result,
	})
}
