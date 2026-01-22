package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
)

type UpdateRutaController struct {
	uc *application.UpdateRutaUseCase
}

func NewUpdateRutaController(uc *application.UpdateRutaUseCase) *UpdateRutaController {
	return &UpdateRutaController{uc}
}

func (ctr *UpdateRutaController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ID inválido"})
		return
	}

	var req struct {
		Nombre      string          `json:"nombre" binding:"required"`
		Descripcion string          `json:"descripcion"`
		JsonRuta    json.RawMessage `json:"json_ruta" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Validar que JsonRuta sea un JSON válido
	if !json.Valid(req.JsonRuta) {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "json_ruta inválido"})
		return
	}

	ruta := &entities.Ruta{
		RutaID:      int32(id),
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
		JsonRuta:    string(req.JsonRuta),
	}

	err = ctr.uc.Run(ruta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	var jsonRuta interface{}
	if err := json.Unmarshal([]byte(ruta.JsonRuta), &jsonRuta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al parsear json_ruta"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"ruta_id":     ruta.RutaID,
		"nombre":      ruta.Nombre,
		"descripcion": ruta.Descripcion,
		"json_ruta":   jsonRuta,
		"eliminado":   ruta.Eliminado,
		"created_at":  ruta.CreatedAt,
	}})
}