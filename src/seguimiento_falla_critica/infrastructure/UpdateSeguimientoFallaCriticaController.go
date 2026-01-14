// UpdateSeguimientoFallaCriticaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/application"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type UpdateSeguimientoFallaCriticaController struct {
	updateUseCase *application.UpdateSeguimientoFallaCriticaUseCase
}

func NewUpdateSeguimientoFallaCriticaController(updateUseCase *application.UpdateSeguimientoFallaCriticaUseCase) *UpdateSeguimientoFallaCriticaController {
	return &UpdateSeguimientoFallaCriticaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateSeguimientoFallaCriticaController) Run(c *gin.Context) {
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
		FallaID    int32  `json:"falla_id" binding:"required"`
		Comentario string `json:"comentario" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	seguimiento := &entities.SeguimientoFallaCritica{
		SeguimientoID: int32(id),
		FallaID:       request.FallaID,
		Comentario:    request.Comentario,
	}

	updatedSeguimiento, err := ctrl.updateUseCase.Run(seguimiento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el seguimiento de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedSeguimiento)
}