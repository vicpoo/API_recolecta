package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
)

type CreateHistorialAsignacionCamionController struct {
	uc *application.SaveHistorialAsignacionCamionUseCase
}

func NewCreateHistorialAsignacionCamionController(uc *application.SaveHistorialAsignacionCamionUseCase) *CreateHistorialAsignacionCamionController {
	return &CreateHistorialAsignacionCamionController{uc: uc}
}

// @Summary      Crear historial de asignación de camión
// @Tags         HistorialAsignacion
// @Accept       json
// @Produce      json
// @Param        body body entities.CreateHistorialAsignacionCamionRequest true "Body"
// @Success      201 {object} entities.HistorialAsignacionCamionResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/historial-asignacion/ [post]
func (ctr *CreateHistorialAsignacionCamionController) Run(ctx *gin.Context) {
	var historial entities.HistorialAsignacionCamion

	if err := ctx.ShouldBindJSON(&historial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	result, err := ctr.uc.Run(&historial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "data": result})
}
