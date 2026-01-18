package application

import (
	"time"

	"github.com/vicpoo/API_recolecta/src/domicilio/domain"
)

type CreateDomicilio struct {
	repo domain.DomicilioRepository
}

func NewCreateDomicilio(repo domain.DomicilioRepository) *CreateDomicilio {
	return &CreateDomicilio{repo}
}

func (uc *CreateDomicilio) Execute(d *domain.Domicilio) error {
	d.Eliminado = false
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return uc.repo.Create(d)
}
