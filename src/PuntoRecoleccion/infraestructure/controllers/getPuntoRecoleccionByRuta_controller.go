package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
)

type GetPuntoRecoleccionByRutaController struct {
	uc *application.GetPuntoRecoleccionByRutaUseCase
}

func NewGetPuntoRecoleccionByRutaController(uc *application.GetPuntoRecoleccionByRutaUseCase) *GetPuntoRecoleccionByRutaController {
	return &GetPuntoRecoleccionByRutaController{uc: uc}
}

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
