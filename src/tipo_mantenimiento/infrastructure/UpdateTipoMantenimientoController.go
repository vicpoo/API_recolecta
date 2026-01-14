// UpdateTipoMantenimientoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain/entities"
)

type UpdateTipoMantenimientoController struct {
	updateUseCase *application.UpdateTipoMantenimientoUseCase
}

func NewUpdateTipoMantenimientoController(updateUseCase *application.UpdateTipoMantenimientoUseCase) *UpdateTipoMantenimientoController {
	return &UpdateTipoMantenimientoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateTipoMantenimientoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		Nombre    string `json:"nombre" binding:"required"`
		Categoria string `json:"categoria" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Crear entidad con solo los campos que se actualizan
	tipoMantenimiento := &entities.TipoMantenimiento{
		ID:        int32(id),
		Nombre:    request.Nombre,
		Categoria: request.Categoria,
	}

	updatedTipoMantenimiento, err := ctrl.updateUseCase.Run(tipoMantenimiento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el tipo de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedTipoMantenimiento)
}