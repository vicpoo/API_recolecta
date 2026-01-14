package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
)

type GetAllRegistroVaciadoController struct {
	uc *application.ListAllRegistroVaciadoUseCase
}

func NewGetAllRegistroVaciadoController(uc *application.ListAllRegistroVaciadoUseCase) *GetAllRegistroVaciadoController {
	return &GetAllRegistroVaciadoController{uc: uc}
}

func (c *GetAllRegistroVaciadoController) Run(ctx *gin.Context) {
	result, err := c.uc.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
