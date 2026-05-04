package application

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/domain"
	"github.com/vicpoo/API_recolecta/src/empleado/domain/entities"
)

type LoginEmpleadoInput struct {
	MailOrUsername string `json:"mail_or_username"`
	Password       string `json:"password"`
}

type LoginEmpleadoOutput struct {
	Empleado *entities.Empleado `json:"empleado"`
	Token    string             `json:"token"`
}

type LoginEmpleado struct {
	repo domain.EmpleadoRepository
}

func NewLoginEmpleado(repo domain.EmpleadoRepository) *LoginEmpleado {
	return &LoginEmpleado{repo: repo}
}

func (uc *LoginEmpleado) Execute(ctx context.Context, in LoginEmpleadoInput) (*LoginEmpleadoOutput, error) {
	credential := strings.TrimSpace(strings.ToLower(in.MailOrUsername))
	password := strings.TrimSpace(in.Password)

	if credential == "" {
		return nil, errors.New("mail_or_username es requerido")
	}

	if password == "" {
		return nil, errors.New("password es requerido")
	}

	empleado, err := uc.repo.FindByMailOrUsername(ctx, credential)
	if err != nil {
		return nil, err
	}

	if empleado == nil {
		return nil, errors.New("credenciales inválidas")
	}

	if empleado.Desactivado {
		return nil, errors.New("empleado desactivado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(empleado.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	token, err := core.GenerateToken(empleado.ID, empleado.RolID)
	if err != nil {
		return nil, err
	}

	return &LoginEmpleadoOutput{
		Empleado: empleado,
		Token:    token,
	}, nil
}
