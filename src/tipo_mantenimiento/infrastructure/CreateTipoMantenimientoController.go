// CreateTipoMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain/entities"
)

type CreateTipoMantenimientoController struct {
	createUseCase *application.CreateTipoMantenimientoUseCase
}

func NewCreateTipoMantenimientoController(createUseCase *application.CreateTipoMantenimientoUseCase) *CreateTipoMantenimientoController {
	return &CreateTipoMantenimientoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateTipoMantenimientoController) Run(c *gin.Context) {
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

	// Crear el tipo de mantenimiento usando el constructor corregido
	tipoMantenimiento := entities.NewTipoMantenimiento(
		request.Nombre,
		request.Categoria,
	)

	createdTipoMantenimiento, err := ctrl.createUseCase.Run(tipoMantenimiento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el tipo de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdTipoMantenimiento)
}