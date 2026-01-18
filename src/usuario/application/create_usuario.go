package application

import (
	"time"

	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"golang.org/x/crypto/bcrypt"
)

type CreateUsuario struct {
	repo domain.UsuarioRepository
}

func NewCreateUsuario(repo domain.UsuarioRepository) *CreateUsuario {
	return &CreateUsuario{repo}
}

func (uc *CreateUsuario) Execute(u *domain.Usuario) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Eliminado = false

	return uc.repo.Create(u)
}
