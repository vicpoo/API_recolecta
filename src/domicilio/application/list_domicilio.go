package application

import "github.com/vicpoo/API_recolecta/src/domicilio/domain"

type ListDomicilios struct {
	repo domain.DomicilioRepository
}

func NewListDomicilios(repo domain.DomicilioRepository) *ListDomicilios {
	return &ListDomicilios{repo}
}

func (uc *ListDomicilios) Execute(usuarioID int) ([]domain.Domicilio, error) {
	return uc.repo.GetAllByUsuario(usuarioID)
}
