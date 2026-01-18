package application

import(
	"errors"
	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsuario struct {
	repo domain.UsuarioRepository
}

func NewLoginUsuario(repo domain.UsuarioRepository) *LoginUsuario {
	return &LoginUsuario{repo}
}
