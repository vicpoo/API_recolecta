package application

import "github.com/vicpoo/API_recolecta/src/domicilio/domain"

type GetDomicilio struct {
	repo domain.DomicilioRepository
}

func NewGetDomicilio(repo domain.DomicilioRepository) *GetDomicilio {
	return &GetDomicilio{repo}
}

func (uc *GetDomicilio) Execute(id int) (*domain.Domicilio, error) {
	return uc.repo.GetByID(id)
}
