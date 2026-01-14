package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
)

type CerrarAsignacionActivaCamionController struct {
	uc *application.CerrarAsignacionActivaCamionUseCase
}

func NewCerrarAsignacionActivaCamionController(uc *application.CerrarAsignacionActivaCamionUseCase) *CerrarAsignacionActivaCamionController {
	return &CerrarAsignacionActivaCamionController{uc: uc}
}

func (ctr *CerrarAsignacionActivaCamionController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("camionId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "camionId inválido"})
		return
	}

	if err := ctr.uc.Run(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "asignación activa del camión cerrada"})
}
