package application

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
)

type UpdateUserInput struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Alias    *string `json:"alias,omitempty"`
	Telefono *string `json:"telefono,omitempty"`
	Email    string  `json:"email"`
	Password *string `json:"password,omitempty"`
	RolID    int     `json:"role_id"`
}

type UpdateUser struct {
	repo domain.UsuarioRepository
}

func NewUpdateUser(repo domain.UsuarioRepository) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (uc *UpdateUser) Execute(ctx context.Context, in UpdateUserInput) error {
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))

	u, err := uc.repo.GetByID(ctx, in.ID)
	if err != nil {
		return err
	}
	if u == nil {
		return nil
	}

	u.Nombre = strings.TrimSpace(in.Nombre)
	u.Alias = in.Alias
	u.Telefono = in.Telefono
	u.Email = in.Email
	u.RolID = in.RolID

	// Si se proporciona una nueva contraseña, hashearla
	if in.Password != nil && *in.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(*in.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PasswordHash = string(hash)
	}

	return uc.repo.Update(ctx, u)
}
