package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
)

type GetAllRutaController struct {
	uc *application.ListAllRutaUseCase
}

func NewGetAllRutaController(uc *application.ListAllRutaUseCase) *GetAllRutaController {
	return &GetAllRutaController{uc}
}

func (ctr *GetAllRutaController) Run(ctx *gin.Context) {
	rutas, err := ctr.uc.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Convertir cada json_ruta de string a objeto
	var rutasResponse []gin.H
	for _, ruta := range rutas {
		var jsonRuta interface{}
		if err := json.Unmarshal([]byte(ruta.JsonRuta), &jsonRuta); err != nil {
			// Si hay error al parsear, devolver el string original
			jsonRuta = ruta.JsonRuta
		}

		rutasResponse = append(rutasResponse, gin.H{
			"ruta_id":     ruta.RutaID,
			"nombre":      ruta.Nombre,
			"descripcion": ruta.Descripcion,
			"json_ruta":   jsonRuta,
			"eliminado":   ruta.Eliminado,
			"created_at":  ruta.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": rutasResponse})
}