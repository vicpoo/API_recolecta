package application

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"github.com/vicpoo/API_recolecta/src/usuario/domain/entities"
)

type SaveUserInput struct {
	Nombre   string
	Email    string
	Password string
	RolID    int
}

type SaveUser struct {
	repo domain.UsuarioRepository
}

func NewSaveUser(repo domain.UsuarioRepository) *SaveUser {
	return &SaveUser{repo: repo}
}

func (uc *SaveUser) Execute(ctx context.Context, in SaveUserInput) (int, error) {
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	u := &entities.Usuario{
		Nombre:       strings.TrimSpace(in.Nombre),
		Email:        in.Email,
		PasswordHash: string(hash),
		RolID:        in.RolID,
	}

	return uc.repo.Create(ctx, u)
}
