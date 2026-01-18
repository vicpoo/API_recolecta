package application

import (
	"errors"

	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsuario struct {
	repo domain.UsuarioRepository
}

func NewLoginUsuario(repo domain.UsuarioRepository) *LoginUsuario {
	return &LoginUsuario{repo}
}

func (uc *LoginUsuario) Execute(email, password string) (string, error) {
	u, err := uc.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	return core.GenerateToken(u.UserID, u.RoleID)
}
