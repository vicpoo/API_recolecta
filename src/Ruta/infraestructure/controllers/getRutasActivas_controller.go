package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
)

type GetRutaActivasController struct {
	uc *application.GetRutaActivasUseCase
}

func NewGetRutaActivasController(uc *application.GetRutaActivasUseCase) *GetRutaActivasController {
	return &GetRutaActivasController{uc}
}

func (ctr *GetRutaActivasController) Run(ctx *gin.Context) {
	rutas, err := ctr.uc.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": rutas})
}
