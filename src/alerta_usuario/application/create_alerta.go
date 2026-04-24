package application

import "github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"

type CreateAlerta struct {
	repo domain.AlertaUsuarioRepository
}

func NewCreateAlerta(repo domain.AlertaUsuarioRepository) *CreateAlerta {
	return &CreateAlerta{repo: repo}
}

func (uc *CreateAlerta) Execute(a *domain.AlertaUsuario) error {
	return uc.repo.Create(a)
}
