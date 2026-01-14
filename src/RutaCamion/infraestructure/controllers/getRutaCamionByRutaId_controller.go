package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/application"
)

type GetRutaCamionByRutaIDController struct {
	uc *application.GetRutaCamionByRutaIDUseCase
}

func NewGetRutaCamionByRutaIDController(
	uc *application.GetRutaCamionByRutaIDUseCase,
) *GetRutaCamionByRutaIDController {
	return &GetRutaCamionByRutaIDController{uc}
}

func (c *GetRutaCamionByRutaIDController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("ruta_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ruta ID inválido"})
		return
	}

	result, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
