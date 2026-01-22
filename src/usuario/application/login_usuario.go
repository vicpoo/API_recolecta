package application

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"github.com/vicpoo/API_recolecta/src/usuario/domain/entities"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginUser struct {
	repo domain.UsuarioRepository
}

func NewLoginUser(repo domain.UsuarioRepository) *LoginUser {
	return &LoginUser{repo: repo}
}

// Devuelve el usuario si las credenciales son válidas.
// Si tú ya generas JWT en otro lado, aquí es donde lo “enchufas”.
func (uc *LoginUser) Execute(ctx context.Context, in LoginInput) (*entities.Usuario, bool, error) {
	email := strings.TrimSpace(strings.ToLower(in.Email))

	u, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, false, err
	}
	if u == nil {
		return nil, false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)); err != nil {
		return nil, false, nil
	}

	return u, true, nil
}
