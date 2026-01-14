package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
)

type DeleteHistorialAsignacionCamionController struct {
	uc *application.DeleteHistorialAsignacionCamionUseCase
}

func NewDeleteHistorialAsignacionCamionController(uc *application.DeleteHistorialAsignacionCamionUseCase) *DeleteHistorialAsignacionCamionController {
	return &DeleteHistorialAsignacionCamionController{uc: uc}
}

func (ctr *DeleteHistorialAsignacionCamionController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "id inválido"})
		return
	}

	if err := ctr.uc.Run(int32(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "registro eliminado"})
}
