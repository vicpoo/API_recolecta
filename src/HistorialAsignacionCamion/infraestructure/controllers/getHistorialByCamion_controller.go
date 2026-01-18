package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
)

type GetHistorialByCamionController struct {
	uc *application.GetHistorialByCamionUseCase
}

func NewGetHistorialByCamionController(uc *application.GetHistorialByCamionUseCase) *GetHistorialByCamionController {
	return &GetHistorialByCamionController{uc: uc}
}

func (ctr *GetHistorialByCamionController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("camionId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "camionId inválido"})
		return
	}

	data, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
