//UpdateAlertaMantenimientoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type UpdateAlertaMantenimientoController struct {
	updateUseCase *application.UpdateAlertaMantenimientoUseCase
}

func NewUpdateAlertaMantenimientoController(updateUseCase *application.UpdateAlertaMantenimientoUseCase) *UpdateAlertaMantenimientoController {
	return &UpdateAlertaMantenimientoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateAlertaMantenimientoController) Run(c *gin.Context) {
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
		CamionID            int32  `json:"camion_id" binding:"required"`
		TipoMantenimientoID int32  `json:"tipo_mantenimiento_id" binding:"required"`
		Descripcion         string `json:"descripcion" binding:"required"`
		Observaciones       string `json:"observaciones"`
		Atendido            bool   `json:"atendido"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Primero obtenemos la alerta existente para preservar la fecha de creación
	alerta, err := ctrl.updateUseCase.Run(&entities.AlertaMantenimiento{
		AlertaID:            int32(id),
		CamionID:            request.CamionID,
		TipoMantenimientoID: request.TipoMantenimientoID,
		Descripcion:         request.Descripcion,
		Observaciones:       request.Observaciones,
		Atendido:            request.Atendido,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la alerta de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alerta)
}