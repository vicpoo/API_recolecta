package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/entities"
	"github.com/vicpoo/API_recolecta/src/ciudadano/domain/ports"
)

type PostgresCiudadanoRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCiudadanoRepository(db *pgxpool.Pool) ports.CiudadanoPostgresRepository {
	return &PostgresCiudadanoRepository{db: db}
}

func (r *PostgresCiudadanoRepository) Create(ctx context.Context, u *entities.CiudadanoPostgres) (int, error) {
	const q = `
		INSERT INTO ciudadano (email, alias, password, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, q, u.Email, u.Alias, u.PasswordHash).Scan(&id)
	return id, err
}

func (r *PostgresCiudadanoRepository) FindByEmail(ctx context.Context, email string) (*entities.CiudadanoPostgres, error) {
	const q = `
		SELECT id, email, alias, password, created_at, updated_at
		FROM ciudadano
		WHERE email = $1
	`
	var u entities.CiudadanoPostgres
	err := r.db.QueryRow(ctx, q, email).Scan(
		&u.ID, &u.Email, &u.Alias, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *PostgresCiudadanoRepository) FindByID(ctx context.Context, id int) (*entities.CiudadanoPostgres, error) {
	const q = `
		SELECT id, email, alias, password, created_at, updated_at
		FROM ciudadano
		WHERE id = $1
	`
	var u entities.CiudadanoPostgres
	err := r.db.QueryRow(ctx, q, id).Scan(
		&u.ID, &u.Email, &u.Alias, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
