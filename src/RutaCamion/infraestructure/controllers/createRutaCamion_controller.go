package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/application"
	"github.com/vicpoo/API_recolecta/src/RutaCamion/domain/entities"
)

type CreateRutaCamionController struct {
	uc *application.SaveRutaCamionUseCase
}

func NewCreateRutaCamionController(
	uc *application.SaveRutaCamionUseCase,
) *CreateRutaCamionController {
	return &CreateRutaCamionController{uc}
}

func (c *CreateRutaCamionController) Run(ctx *gin.Context) {
	var rutaCamion entities.RutaCamion

	if err := ctx.ShouldBindJSON(&rutaCamion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(&rutaCamion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
