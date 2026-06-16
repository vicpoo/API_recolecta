package application_domicilio

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
)

type CreateDomicilioInput struct {
	CiudadanoID int     `json:"ciudadano_id"`
	ColoniaID   int     `json:"colonia_id"`
	Alias       string  `json:"alias"`
	Calle       string  `json:"calle"`
	Numero      string  `json:"numero"`
	Referencia  *string `json:"referencia,omitempty"`
}

type CreateDomicilio struct {
	repo domain.DomicilioRepository
}

func NewCreateDomicilio(repo domain.DomicilioRepository) *CreateDomicilio {
	return &CreateDomicilio{repo: repo}
}

func (uc *CreateDomicilio) Execute(ctx context.Context, in CreateDomicilioInput) (int, error) {
	in.Alias = strings.TrimSpace(in.Alias)
	in.Calle = strings.TrimSpace(in.Calle)
	in.Numero = strings.TrimSpace(in.Numero)

	if in.CiudadanoID <= 0 {
		return 0, errors.New("ciudadano_id es requerido")
	}
	if in.ColoniaID <= 0 {
		return 0, errors.New("colonia_id es requerido")
	}
	if in.Alias == "" {
		return 0, errors.New("alias es requerido")
	}
	if in.Calle == "" {
		return 0, errors.New("calle es requerida")
	}
	if in.Numero == "" {
		return 0, errors.New("numero es requerido")
	}

	existingByAlias, err := uc.repo.FindByAlias(ctx, in.Alias)
	if err != nil {
		return 0, err
	}
	if existingByAlias != nil {
		return 0, errors.New("el alias del domicilio ya está registrado")
	}

	d := &entities.Domicilio{
		CiudadanoID: in.CiudadanoID,
		ColoniaID:   in.ColoniaID,
		Alias:       in.Alias,
		Calle:       in.Calle,
		Numero:      in.Numero,
		Referencia:  in.Referencia,
		CreatedAt:   time.Now(),
	}

	return uc.repo.Create(ctx, d)
}