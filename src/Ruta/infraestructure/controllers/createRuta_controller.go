package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
)

type CreateRutaController struct {
	uc *application.CreateRutaUseCase
}

func NewCreateRutaController(uc *application.CreateRutaUseCase) *CreateRutaController {
	return &CreateRutaController{uc}
}

func (ctr *CreateRutaController) Run(ctx *gin.Context) {
	var ruta entities.Ruta

	if err := ctx.ShouldBindJSON(&ruta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	err := ctr.uc.Run(&ruta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "data": ruta})
}
