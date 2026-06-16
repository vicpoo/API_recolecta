package application_domicilio

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
)

type ListDomicilios struct {
	repo domain.DomicilioRepository
}

func NewListDomicilios(repo domain.DomicilioRepository) *ListDomicilios {
	return &ListDomicilios{repo: repo}
}

func (uc *ListDomicilios) Execute(ctx context.Context) ([]entities.Domicilio, error) {
	return uc.repo.List(ctx)
}

func (uc *ListDomicilios) ExecuteByCiudadanoID(ctx context.Context, ciudadanoID int) ([]entities.Domicilio, error) {
	return uc.repo.ListByCiudadanoID(ctx, ciudadanoID)
}