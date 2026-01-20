package application

import "time"

import "github.com/vicpoo/API_recolecta/src/domicilio/domain"

type UpdateDomicilio struct {
	repo domain.DomicilioRepository
}

func NewUpdateDomicilio(repo domain.DomicilioRepository) *UpdateDomicilio {
	return &UpdateDomicilio{repo}
}

func (uc *UpdateDomicilio) Execute(d *domain.Domicilio) error {
	d.UpdatedAt = time.Now()
	return uc.repo.Update(d)
}
