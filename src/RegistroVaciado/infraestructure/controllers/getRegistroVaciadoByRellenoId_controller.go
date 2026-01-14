package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
)

type GetRegistroVaciadoByRellenoIDController struct {
	uc *application.GetRegistroVaciadoByRellenoIDUseCase
}

func NewGetRegistroVaciadoByRellenoIDController(
	uc *application.GetRegistroVaciadoByRellenoIDUseCase,
) *GetRegistroVaciadoByRellenoIDController {
	return &GetRegistroVaciadoByRellenoIDController{uc: uc}
}

func (c *GetRegistroVaciadoByRellenoIDController) Run(ctx *gin.Context) {
	rellenoID, err := strconv.Atoi(ctx.Param("relleno_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "relleno_id inválido"})
		return
	}

	result, err := c.uc.Execute(int32(rellenoID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
