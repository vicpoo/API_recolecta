package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type GetPuntoRecoleccionByRutaController struct {
	uc *application.GetPuntoRecoleccionByRutaUseCase
}

func NewGetPuntoRecoleccionByRutaController(uc *application.GetPuntoRecoleccionByRutaUseCase) *GetPuntoRecoleccionByRutaController {
	return &GetPuntoRecoleccionByRutaController{uc: uc}
}

// @Summary      Puntos de recolección por ruta
// @Tags         PuntoRecoleccion
// @Produce      json
// @Param        rutaId path int true "ID ruta"
// @Success      200 {array} map[string]interface{}
// @Router       /api/puntos-recoleccion/ruta/{rutaId} [get]
func (c *GetPuntoRecoleccionByRutaController) Run(ctx *gin.Context) {
	rutaStr := ctx.Param("rutaId")
	rutaId, err := strconv.Atoi(rutaStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "rutaId inválido"})
		return
	}

	result, err := c.uc.Execute(int32(rutaId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
