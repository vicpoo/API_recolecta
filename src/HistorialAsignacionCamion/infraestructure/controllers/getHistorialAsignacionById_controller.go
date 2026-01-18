package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
)

type GetHistorialAsignacionByIdController struct {
	uc *application.GetHistorialAsignacionCamionByIdUseCase
}

func NewGetHistorialAsignacionByIdController(uc *application.GetHistorialAsignacionCamionByIdUseCase) *GetHistorialAsignacionByIdController {
	return &GetHistorialAsignacionByIdController{uc: uc}
}

func (ctr *GetHistorialAsignacionByIdController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "id inválido"})
		return
	}

	result, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}
