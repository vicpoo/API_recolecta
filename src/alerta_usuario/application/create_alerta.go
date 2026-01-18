package application

import (
	"time"

	"github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"
)

type CreateAlerta struct {
	repo domain.AlertaUsuarioRepository
}

func NewCreateAlerta(repo domain.AlertaUsuarioRepository) *CreateAlerta {
	return &CreateAlerta{repo}
}

func (uc *CreateAlerta) Execute(a *domain.AlertaUsuario) error {
	a.Leida = false
	a.CreatedAt = time.Now()
	return uc.repo.Create(a)
}
