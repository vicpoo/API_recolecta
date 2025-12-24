package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
)

type CreateCamionController struct {
	uc *application.SaveCamionUseCase
}

func NewCreateCamionController(uc *application.SaveCamionUseCase) *CreateCamionController {
	return &CreateCamionController{
		uc: uc,
	}
}

func (ctr *CreateCamionController) Run(ctx *gin.Context) {
	var camion entities.Camion

	if err := ctx.ShouldBindJSON(&camion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "datos inválidos",
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	result, err := ctr.uc.Run(&camion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "error al crear el camión",
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "camión creado correctamente",
		"data":    result,
		"code":    http.StatusCreated,
	})
}
