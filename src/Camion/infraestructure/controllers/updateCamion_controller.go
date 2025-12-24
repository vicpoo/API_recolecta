package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
)

type UpdateCamionController struct {
	uc *application.UpdateCamionUseCase
}

func NewUpdateCamionController(uc *application.UpdateCamionUseCase) *UpdateCamionController {
	return &UpdateCamionController{
		uc: uc,
	}
}

func (ctr *UpdateCamionController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id inválido",
			"code":    http.StatusBadRequest,
		})
		return
	}

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

	result, err := ctr.uc.Run(int32(id), &camion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "camión actualizado correctamente",
		"data":    result,
		"code":    http.StatusOK,
	})
}
