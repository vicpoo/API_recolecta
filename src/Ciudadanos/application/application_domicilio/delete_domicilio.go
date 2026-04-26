package application_domicilio

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
)

type DeleteDomicilio struct {
	repo domain.DomicilioRepository
}

func NewDeleteDomicilio(repo domain.DomicilioRepository) *DeleteDomicilio {
	return &DeleteDomicilio{repo: repo}
}

func (uc *DeleteDomicilio) Execute(ctx context.Context, id int, ciudadanoID int) error {
	return uc.repo.DeleteByCiudadano(ctx, id, ciudadanoID)
}