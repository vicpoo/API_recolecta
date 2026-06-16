package controller_domicilio

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_domicilio"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DomicilioController struct {
	create *application_domicilio.CreateDomicilio
	get    *application_domicilio.GetDomicilio
	update *application_domicilio.UpdateDomicilio
	list   *application_domicilio.ListDomicilios
	delete *application_domicilio.DeleteDomicilio
}

func NewDomicilioController(
	create *application_domicilio.CreateDomicilio,
	get *application_domicilio.GetDomicilio,
	update *application_domicilio.UpdateDomicilio,
	list *application_domicilio.ListDomicilios,
	delete *application_domicilio.DeleteDomicilio,
) *DomicilioController {
	return &DomicilioController{
		create: create,
		get:    get,
		update: update,
		list:   list,
		delete: delete,
	}
}

func (c *DomicilioController) Create(ctx *gin.Context) {
	var body application_domicilio.CreateDomicilioInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		core.RespondValidationError(ctx, "Datos de domicilio inválidos", map[string]string{"error": err.Error()})
		return
	}

	if body.CiudadanoID == 0 {
		body.CiudadanoID = ctx.GetInt("user_id")
	}

	id, err := c.create.Execute(ctx.Request.Context(), body)
	if err != nil {
		core.RespondInternalServerError(ctx, "Error al crear domicilio", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "domicilio creado correctamente",
		"id":      id,
	})
}

func (c *DomicilioController) List(ctx *gin.Context) {
	ciudadanoID := ctx.GetInt("user_id")

	domicilios, err := c.list.ExecuteByCiudadanoID(ctx.Request.Context(), ciudadanoID)
	if err != nil {
		core.RespondInternalServerError(ctx, "Error al listar domicilios", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": domicilios,
	})
}

func (c *DomicilioController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de domicilio inválido")
		return
	}

	domicilio, err := c.get.Execute(ctx.Request.Context(), id)
	if err != nil {
		core.RespondInternalServerError(ctx, "Error al obtener domicilio", err)
		return
	}

	if domicilio == nil {
		core.RespondError(ctx, http.StatusNotFound, core.ErrCodeNotFound, "Domicilio no encontrado", nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": domicilio,
	})
}

func (c *DomicilioController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de domicilio inválido")
		return
	}

	var body application_domicilio.UpdateDomicilioInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		core.RespondValidationError(ctx, "Datos de domicilio inválidos", map[string]string{"error": err.Error()})
		return
	}

	body.ID = id

	if err := c.update.Execute(ctx.Request.Context(), body); err != nil {
		core.RespondInternalServerError(ctx, "Error al actualizar domicilio", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "domicilio actualizado correctamente",
	})
}

func (c *DomicilioController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de domicilio inválido")
		return
	}

	ciudadanoID := ctx.GetInt("user_id")

	if err := c.delete.Execute(ctx.Request.Context(), id, ciudadanoID); err != nil {
		core.RespondError(ctx, http.StatusForbidden, core.ErrCodeForbidden, "No autorizado para eliminar este domicilio", map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "domicilio eliminado correctamente",
	})
}
