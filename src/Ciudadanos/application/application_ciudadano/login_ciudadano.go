package  application_ciudadano

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
)

type LoginCiudadanoInput struct {
	EmailOrAlias string `json:"email_or_alias"`
	Password     string `json:"password"`
}

type LoginCiudadano struct {
	repo domain.CiudadanoRepository
}

func NewLoginCiudadano(repo domain.CiudadanoRepository) *LoginCiudadano {
	return &LoginCiudadano{repo: repo}
}

func (uc *LoginCiudadano) Execute(ctx context.Context, in LoginCiudadanoInput) (*entities.Ciudadano, bool, error) {
	credential := strings.TrimSpace(strings.ToLower(in.EmailOrAlias))
	password := strings.TrimSpace(in.Password)

	ciudadano, err := uc.repo.FindByEmailOrAlias(ctx, credential)
	if err != nil {
		return nil, false, err
	}
	if ciudadano == nil {
		return nil, false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(ciudadano.Password), []byte(password)); err != nil {
		return nil, false, nil
	}

	return ciudadano, true, nil
}