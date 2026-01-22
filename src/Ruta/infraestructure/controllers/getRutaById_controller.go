package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
)

type GetRutaByIdController struct {
	uc *application.GetRutaByIdUseCase
}

func NewGetRutaByIdController(uc *application.GetRutaByIdUseCase) *GetRutaByIdController {
	return &GetRutaByIdController{uc: uc}
}

func (ctr *GetRutaByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID inválido",
		})
		return
	}

	ruta, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Ruta no encontrada",
		})
		return
	}

	// Convertir json_ruta de string a objeto JSON
	var jsonRuta interface{}
	if err := json.Unmarshal([]byte(ruta.JsonRuta), &jsonRuta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error al parsear json_ruta",
		})
		return
	}

	// Crear respuesta con json_ruta como objeto
	response := gin.H{
		"success": true,
		"data": gin.H{
			"ruta_id":     ruta.RutaID,
			"nombre":      ruta.Nombre,
			"descripcion": ruta.Descripcion,
			"json_ruta":   jsonRuta, // Ahora es un objeto, no string
			"eliminado":   ruta.Eliminado,
			"created_at":  ruta.CreatedAt,
		},
	}

	ctx.JSON(http.StatusOK, response)
}