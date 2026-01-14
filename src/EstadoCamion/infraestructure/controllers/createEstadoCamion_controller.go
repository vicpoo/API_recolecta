package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/application"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
)

type CreateEstadoCamionController struct {
	uc *application.SaveEstadoCamionUseCase
}

func NewCreateEstadoCamionController(
	uc *application.SaveEstadoCamionUseCase,
) *CreateEstadoCamionController {
	return &CreateEstadoCamionController{
		uc: uc,
	}
}

func (ctr *CreateEstadoCamionController) Run(ctx *gin.Context) {
	var estadoCamion entities.EstadoCamion

	if err := ctx.ShouldBindJSON(&estadoCamion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "body inválido",
			"error":   err.Error(),
		})
		return
	}

	estadoCamionSaved, err := ctr.uc.Run(&estadoCamion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    estadoCamionSaved,
	})
}
