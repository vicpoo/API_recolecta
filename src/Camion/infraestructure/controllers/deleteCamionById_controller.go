package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
)

type DeleteCamionController struct {
	uc *application.DeleteCamionUseCase
}

func NewDeleteCamionController(uc *application.DeleteCamionUseCase) *DeleteCamionController {
	return &DeleteCamionController{
		uc: uc,
	}
}

func (ctr *DeleteCamionController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id inválido",
		})
		return
	}

	err = ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "camión eliminado correctamente",
	})
}
