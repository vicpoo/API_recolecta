package application

import (
	"context"

	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/ports"
)

type UpdatePuntoRecoleccionUseCase struct {
	repo ports.IPuntoRecoleccion
}

func NewUpdatePuntoRecoleccionUseCase(repo ports.IPuntoRecoleccion) *UpdatePuntoRecoleccionUseCase {
	return &UpdatePuntoRecoleccionUseCase{repo: repo}
}

func (uc *UpdatePuntoRecoleccionUseCase) Execute(ctx context.Context, tenantID int, id int32, p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error) {
	return uc.repo.Update(ctx, tenantID, id, p)
}
