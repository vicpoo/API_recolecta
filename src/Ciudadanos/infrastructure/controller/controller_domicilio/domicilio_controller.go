package controller_domicilio

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_domicilio"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "json inválido",
			"detail": err.Error(),
		})
		return
	}

	if body.CiudadanoID == 0 {
		body.CiudadanoID = ctx.GetInt("user_id")
	}

	id, err := c.create.Execute(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "domicilio creado correctamente",
		"id":      id,
	})
}

func (c *DomicilioController) List(ctx *gin.Context) {
	domicilios, err := c.list.Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": domicilios,
	})
}

func (c *DomicilioController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	domicilio, err := c.get.Execute(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if domicilio == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "domicilio no encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": domicilio,
	})
}

func (c *DomicilioController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	var body application_domicilio.UpdateDomicilioInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "json inválido",
			"detail": err.Error(),
		})
		return
	}

	body.ID = id

	if err := c.update.Execute(ctx.Request.Context(), body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "domicilio actualizado correctamente",
	})
}

func (c *DomicilioController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	if err := c.delete.Execute(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "domicilio eliminado correctamente",
	})
}