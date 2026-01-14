package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/entities"
)

type CreatePuntoRecoleccionController struct {
	uc *application.SavePuntoRecoleccionUseCase
}

func NewCreatePuntoRecoleccionController(uc *application.SavePuntoRecoleccionUseCase) *CreatePuntoRecoleccionController {
	return &CreatePuntoRecoleccionController{uc: uc}
}

func (c *CreatePuntoRecoleccionController) Run(ctx *gin.Context) {
	var p entities.PuntoRecoleccion

	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(&p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
