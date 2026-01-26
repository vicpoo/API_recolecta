package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/domicilio/domain"
)

type DomicilioRepository struct {
	db *pgxpool.Pool
}

func NewDomicilioRepository(db *pgxpool.Pool) *DomicilioRepository {
	return &DomicilioRepository{db}
}

func (r *DomicilioRepository) Create(d *domain.Domicilio) error {
	query := `
		INSERT INTO domicilio
		(usuario_id, alias, direccion, colonia_id, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		d.UsuarioID,
		d.Alias,
		d.Direccion,
		d.ColoniaID,
		d.CreatedAt,
		d.UpdatedAt,
	)

	return err
}

func (r *DomicilioRepository) GetByID(id int) (*domain.Domicilio, error) {
	query := `
		SELECT domicilio_id, usuario_id, alias, direccion, colonia_id,
		       eliminado, created_at, updated_at
		FROM domicilio
		WHERE domicilio_id = $1 AND eliminado = false
	`

	row := r.db.QueryRow(context.Background(), query, id)

	var d domain.Domicilio
	err := row.Scan(
		&d.DomicilioID,
		&d.UsuarioID,
		&d.Alias,
		&d.Direccion,
		&d.ColoniaID,
		&d.Eliminado,
		&d.CreatedAt,
		&d.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DomicilioRepository)Delete(id int, usuarioID int) error{
	query := `
		UPDATE domicilio
		SET eliminado = true, updated_at = NOW()
		WHERE domicilio_id = $1 AND usuario_id = $2 AND eliminado = false
	`
	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
		usuarioID,
	)
	return err
}

func (r *DomicilioRepository) Update(d *domain.Domicilio) error {
	query := `
		UPDATE domicilio
		SET alias = $1,
		    direccion = $2,
		    colonia_id = $3,
		    updated_at = $4
		WHERE domicilio_id = $5 AND eliminado = false
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		d.Alias,
		d.Direccion,
		d.ColoniaID,
		d.UpdatedAt,
		d.DomicilioID,
	)

	return err
}

func (r *DomicilioRepository) GetAllByUsuario(usuarioID int) ([]domain.Domicilio, error) {
	query := `
		SELECT domicilio_id, usuario_id, alias, direccion, colonia_id,
		       eliminado, created_at, updated_at
		FROM domicilio
		WHERE usuario_id = $1 AND eliminado = false
	`

	rows, err := r.db.Query(context.Background(), query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domicilios []domain.Domicilio

	for rows.Next() {
		var d domain.Domicilio
		if err := rows.Scan(
			&d.DomicilioID,
			&d.UsuarioID,
			&d.Alias,
			&d.Direccion,
			&d.ColoniaID,
			&d.Eliminado,
			&d.CreatedAt,
			&d.UpdatedAt,
		); err != nil {
			return nil, err
		}
		domicilios = append(domicilios, d)
	}

	return domicilios, nil
}

