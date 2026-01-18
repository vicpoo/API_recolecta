package application

import "github.com/vicpoo/API_recolecta/src/domicilio/domain"

type DeleteDomicilio struct {
	repo domain.DomicilioRepository
}

func NewDeleteDomicilio(repo domain.DomicilioRepository) *DeleteDomicilio {
	return &DeleteDomicilio{repo}
}

func (uc *DeleteDomicilio) Execute(id int, usuarioID int) error {
	return uc.repo.Delete(id, usuarioID)
}
