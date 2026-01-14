package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"
)

type CreateRegistroVaciadoController struct {
	uc *application.CreateRegistroVaciadoUseCase
}

func NewCreateRegistroVaciadoController(uc *application.CreateRegistroVaciadoUseCase) *CreateRegistroVaciadoController {
	return &CreateRegistroVaciadoController{uc: uc}
}

func (c *CreateRegistroVaciadoController) Run(ctx *gin.Context) {
	var registro entities.RegistroVaciado

	if err := ctx.ShouldBindJSON(&registro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(&registro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
