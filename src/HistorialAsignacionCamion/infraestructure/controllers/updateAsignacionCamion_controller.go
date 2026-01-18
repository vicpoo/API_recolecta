package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
)

type UpdateHistorialAsignacionCamionController struct {
	uc *application.UpdateHistorialAsignacionCamionUseCase
}

func NewUpdateHistorialAsignacionCamionController(uc *application.UpdateHistorialAsignacionCamionUseCase) *UpdateHistorialAsignacionCamionController {
	return &UpdateHistorialAsignacionCamionController{uc: uc}
}

func (ctr *UpdateHistorialAsignacionCamionController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "id inválido"})
		return
	}

	var historial entities.HistorialAsignacionCamion
	if err := ctx.ShouldBindJSON(&historial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	result, err := ctr.uc.Run(int32(id), &historial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}
