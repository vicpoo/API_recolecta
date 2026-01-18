package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/application"
)

type GetRutaCamionByCamionIDController struct {
	uc *application.GetRutaCamionByCamionIDUseCase
}

func NewGetRutaCamionByCamionIDController(
	uc *application.GetRutaCamionByCamionIDUseCase,
) *GetRutaCamionByCamionIDController {
	return &GetRutaCamionByCamionIDController{uc}
}

func (c *GetRutaCamionByCamionIDController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("camion_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Camión ID inválido"})
		return
	}

	result, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
