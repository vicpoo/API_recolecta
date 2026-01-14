package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
)

type GetHistorialByChoferController struct {
	uc *application.GetHistorialByChoferUseCase
}

func NewGetHistorialByChoferController(uc *application.GetHistorialByChoferUseCase) *GetHistorialByChoferController {
	return &GetHistorialByChoferController{uc: uc}
}

func (ctr *GetHistorialByChoferController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("choferId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "choferId inválido"})
		return
	}

	data, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
