package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
)

type GetRegistroVaciadoByRutaCamionIDController struct {
	uc *application.GetRegistroVaciadoByRutaCamionIDUseCase
}

func NewGetRegistroVaciadoByRutaCamionIDController(
	uc *application.GetRegistroVaciadoByRutaCamionIDUseCase,
) *GetRegistroVaciadoByRutaCamionIDController {
	return &GetRegistroVaciadoByRutaCamionIDController{uc: uc}
}

func (c *GetRegistroVaciadoByRutaCamionIDController) Run(ctx *gin.Context) {
	rutaCamionID, err := strconv.Atoi(ctx.Param("ruta_camion_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ruta_camion_id inválido"})
		return
	}

	result, err := c.uc.Execute(int32(rutaCamionID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
