package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/usuario/application"
)

type AddUsersController struct {
	uc *application.SaveUser
}

func NewAddUsersController(uc *application.SaveUser) *AddUsersController {
	return &AddUsersController{uc: uc}
}

func (c *AddUsersController) Handle(ctx *gin.Context) {
	var body application.SaveUserInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := c.uc.Execute(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}
