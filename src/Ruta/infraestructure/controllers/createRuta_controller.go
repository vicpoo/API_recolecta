package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ruta/application"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
)

type CreateRutaController struct {
	uc *application.CreateRutaUseCase
}

func NewCreateRutaController(uc *application.CreateRutaUseCase) *CreateRutaController {
	return &CreateRutaController{uc}
}

func (ctr *CreateRutaController) Run(ctx *gin.Context) {
	var req struct {
		Nombre      string          `json:"nombre" binding:"required"`
		Descripcion string          `json:"descripcion"`
		JsonRuta    json.RawMessage `json:"json_ruta" binding:"required"`
		CreatedAt   *time.Time      `json:"created_at"` // Opcional, si no viene usa time.Now()
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

	// Si no viene created_at, usar la fecha actual
	createdAt := time.Now()
	if req.CreatedAt != nil {
		createdAt = *req.CreatedAt
	}

	ruta := &entities.Ruta{
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
		JsonRuta:    string(req.JsonRuta),
		CreatedAt:   createdAt,
	}

	err := ctr.uc.Run(ruta)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	var jsonRuta interface{}
	if err := json.Unmarshal([]byte(ruta.JsonRuta), &jsonRuta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al parsear json_ruta"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "data": gin.H{
		"ruta_id":     ruta.RutaID,
		"nombre":      ruta.Nombre,
		"descripcion": ruta.Descripcion,
		"json_ruta":   jsonRuta,
		"eliminado":   ruta.Eliminado,
		"created_at":  ruta.CreatedAt,
	}})
}