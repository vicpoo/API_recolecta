package application_ciudadano

import (
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
)

type CreateCiudadanoInput struct {
	Email    string `json:"email"`
	Alias    string `json:"alias"`
	Password string `json:"password"`
}

type CreateCiudadano struct {
	repo domain.CiudadanoRepository
}

func NewCreateCiudadano(repo domain.CiudadanoRepository) *CreateCiudadano {
	return &CreateCiudadano{repo: repo}
}

func (uc *CreateCiudadano) Execute(ctx context.Context, in CreateCiudadanoInput) (int, error) {
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	in.Alias = strings.TrimSpace(in.Alias)
	in.Password = strings.TrimSpace(in.Password)

	if in.Email == "" {
		return 0, errors.New("email es requerido")
	}
	if in.Alias == "" {
		return 0, errors.New("alias es requerido")
	}
	if in.Password == "" {
		return 0, errors.New("password es requerido")
	}

	existingByEmail, err := uc.repo.FindByEmail(ctx, in.Email)
	if err != nil {
		return 0, err
	}
	if existingByEmail != nil {
		return 0, errors.New("el email ya está registrado")
	}

	existingByAlias, err := uc.repo.FindByAlias(ctx, in.Alias)
	if err != nil {
		return 0, err
	}
	if existingByAlias != nil {
		return 0, errors.New("el alias ya está registrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	ciudadano := &entities.Ciudadano{
		Email:     in.Email,
		Alias:     in.Alias,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	return uc.repo.Create(ctx, ciudadano)
}