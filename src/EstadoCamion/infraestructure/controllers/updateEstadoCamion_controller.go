package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/application"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
)

type UpdateEstadoCamionController struct {
	uc *application.UpdateEstadoCamionUseCase
}

func NewUpdateEstadoCamionController(
	uc *application.UpdateEstadoCamionUseCase,
) *UpdateEstadoCamionController {
	return &UpdateEstadoCamionController{
		uc: uc,
	}
}

func (ctr *UpdateEstadoCamionController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id inválido",
		})
		return
	}

	var estadoCamion entities.EstadoCamion
	if err := ctx.ShouldBindJSON(&estadoCamion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "body inválido",
			"error":   err.Error(),
		})
		return
	}

	estadoCamionUpdated, err := ctr.uc.Run(int32(id), &estadoCamion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    estadoCamionUpdated,
	})
}
