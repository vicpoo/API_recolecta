package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/usuario/application"
)

type LoginUsersController struct {
	uc *application.LoginUser
}

func NewLoginUsersController(uc *application.LoginUser) *LoginUsersController {
	return &LoginUsersController{uc: uc}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *LoginUsersController) Handle(ctx *gin.Context) {
	var body loginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario, valid, err := c.uc.Execute(ctx, application.LoginInput{Email: body.Email, Password: body.Password})
	if err != nil || !valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales inválidas"})
		return
	}

	token, err := core.GenerateToken(usuario.ID, usuario.RolID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al generar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   token,
		"usuario": usuario,
	})
}
