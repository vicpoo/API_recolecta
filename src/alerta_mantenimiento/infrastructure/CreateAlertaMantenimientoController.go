//CreateAlertaMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type CreateAlertaMantenimientoController struct {
	createUseCase *application.CreateAlertaMantenimientoUseCase
}

func NewCreateAlertaMantenimientoController(createUseCase *application.CreateAlertaMantenimientoUseCase) *CreateAlertaMantenimientoController {
	return &CreateAlertaMantenimientoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateAlertaMantenimientoController) Run(c *gin.Context) {
	var request struct {
		CamionID            int32  `json:"camion_id" binding:"required"`
		TipoMantenimientoID int32  `json:"tipo_mantenimiento_id" binding:"required"`
		Descripcion         string `json:"descripcion" binding:"required"`
		Observaciones       string `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	alerta := entities.NewAlertaMantenimiento(
		request.CamionID,
		request.TipoMantenimientoID,
		request.Descripcion,
		request.Observaciones,
	)

	createdAlerta, err := ctrl.createUseCase.Run(alerta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la alerta de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdAlerta)
}