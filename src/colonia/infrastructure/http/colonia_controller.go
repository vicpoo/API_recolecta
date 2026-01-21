package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/colonia/application"
	"github.com/vicpoo/API_recolecta/src/colonia/domain"
)

type ColoniaController struct {
	create *application.CreateColonia
	get    *application.GetColonia
	list   *application.ListColonias
	update *application.UpdateColonia
	delete *application.DeleteColonia
}

func NewColoniaController(
	create *application.CreateColonia,
	get *application.GetColonia,
	list *application.ListColonias,
	update *application.UpdateColonia,
	delete *application.DeleteColonia,
) *ColoniaController {
	return &ColoniaController{create, get, list, update, delete}
}

func (c *ColoniaController) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/colonias")
	{
		group.POST("", c.Create)
		group.GET("", c.List)
		group.GET("/:id", c.GetByID)
		group.PUT("/:id", c.Update)
		group.DELETE("/:id", c.Delete)
	}
}

func (c *ColoniaController) Create(ctx *gin.Context) {
	var body domain.Colonia
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.create.Execute(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, body)
}

func (c *ColoniaController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	colonia, err := c.get.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "colonia no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, colonia)
}

func (c *ColoniaController) List(ctx *gin.Context) {
	colonias, err := c.list.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, colonias)
}

func (c *ColoniaController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var body domain.Colonia
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body.ColoniaID = id

	if err := c.update.Execute(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ColoniaController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.delete.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
