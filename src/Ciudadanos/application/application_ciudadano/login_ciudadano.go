package application_ciudadano

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type LoginCiudadanoInput struct {
	EmailOrAlias string `json:"email_or_alias"`
	Password     string `json:"password"`
}

type LoginCiudadanoOutput struct {
	Ciudadano *entities.Ciudadano `json:"ciudadano"`
	Token     string              `json:"token"`
}

type LoginCiudadano struct {
	repo domain.CiudadanoRepository
}

func NewLoginCiudadano(repo domain.CiudadanoRepository) *LoginCiudadano {
	return &LoginCiudadano{repo: repo}
}

func (uc *LoginCiudadano) Execute(ctx context.Context, in LoginCiudadanoInput) (*LoginCiudadanoOutput, error) {
	credential := strings.TrimSpace(strings.ToLower(in.EmailOrAlias))
	password := strings.TrimSpace(in.Password)

	if credential == "" {
		return nil, errors.New("email_or_alias es requerido")
	}

	if password == "" {
		return nil, errors.New("password es requerido")
	}

	ciudadano, err := uc.repo.FindByEmailOrAlias(ctx, credential)
	if err != nil {
		return nil, err
	}

	if ciudadano == nil {
		return nil, errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(ciudadano.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	token, err := core.GenerateToken(ciudadano.ID, 0)
	if err != nil {
		return nil, err
	}

	return &LoginCiudadanoOutput{
		Ciudadano: ciudadano,
		Token:     token,
	}, nil
}