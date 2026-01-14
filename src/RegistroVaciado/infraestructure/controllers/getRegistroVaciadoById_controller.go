package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
)

type GetRegistroVaciadoByIDController struct {
	uc *application.GetRegistroVaciadoByIDUseCase
}

func NewGetRegistroVaciadoByIDController(uc *application.GetRegistroVaciadoByIDUseCase) *GetRegistroVaciadoByIDController {
	return &GetRegistroVaciadoByIDController{uc: uc}
}

func (c *GetRegistroVaciadoByIDController) Run(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	result, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
