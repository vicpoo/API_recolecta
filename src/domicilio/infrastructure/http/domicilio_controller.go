package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/domicilio/application"
	"github.com/vicpoo/API_recolecta/src/domicilio/domain"
)

type DomicilioController struct {
	create *application.CreateDomicilio
	get    *application.GetDomicilio
	update *application.UpdateDomicilio
}

func NewDomicilioController(
	create *application.CreateDomicilio,
	get *application.GetDomicilio,
	update *application.UpdateDomicilio,
) *DomicilioController {
	return &DomicilioController{create, get, update}
}

func (c *DomicilioController) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/domicilios")
	{
		group.POST("", c.Create)
		group.GET("/:id", c.GetByID)
		group.PUT("/:id", c.Update)
	}
}

func (c *DomicilioController) Create(ctx *gin.Context) {
	var body domain.Domicilio
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.create.Execute(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *DomicilioController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	domicilio, err := c.get.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "domicilio no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, domicilio)
}

func (c *DomicilioController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var body domain.Domicilio
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body.DomicilioID = id

	if err := c.update.Execute(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
