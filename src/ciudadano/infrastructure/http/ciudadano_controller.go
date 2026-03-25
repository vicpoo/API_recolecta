package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/ciudadano/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CiudadanoController struct {
	registerUC      *application.RegisterCiudadanoUseCase
	updateCoordUC   *application.UpdateCoordinatesUseCase
}

func NewCiudadanoController(
	registerUC *application.RegisterCiudadanoUseCase,
	updateCoordUC *application.UpdateCoordinatesUseCase,
) *CiudadanoController {
	return &CiudadanoController{registerUC: registerUC, updateCoordUC: updateCoordUC}
}

func (ctrl *CiudadanoController) RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api/ciudadanos")
	{
		public.POST("/register", ctrl.Register)
	}

	protected := r.Group("/api/ciudadanos", core.JWTAuthMiddleware())
	{
		protected.POST("/coordinates", ctrl.UpdateCoordinates)
	}
}

func (ctrl *CiudadanoController) Register(ctx *gin.Context) {
	var body application.RegisterCiudadanoRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := ctrl.registerUC.Execute(ctx, body)
	if err != nil {
		if err.Error() == "email ya registrado" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := core.GenerateToken(id, core.CIUDADANO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al generar token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id, "token": token})
}

type updateCoordinatesBody struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func (ctrl *CiudadanoController) UpdateCoordinates(ctx *gin.Context) {
	var body updateCoordinatesBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetInt("user_id")

	err := ctrl.updateCoordUC.Execute(ctx, application.UpdateCoordinatesRequest{
		UserID:    userID,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	})
	if err != nil {
		if err.Error() == "ciudadano no encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
