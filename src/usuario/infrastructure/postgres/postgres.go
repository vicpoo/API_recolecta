package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/usuario/domain/entities"
)

type UsuarioPostgresRepository struct {
	db *pgxpool.Pool
}

func NewUsuarioPostgresRepository(db *pgxpool.Pool) *UsuarioPostgresRepository {
	return &UsuarioPostgresRepository{db: db}
}

func (r *UsuarioPostgresRepository) Create(ctx context.Context, u *entities.Usuario) (int, error) {
	const q = `
		INSERT INTO usuario (nombre, email, password, role_id, alias, telefono, residencia_id, eliminado, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING user_id
	`
	var id int
	err := r.db.QueryRow(ctx, q, u.Nombre, u.Email, u.PasswordHash, u.RolID, u.Alias, u.Telefono, u.ResidenciaID, u.Eliminado).Scan(&id)
	return id, err
}

func (r *UsuarioPostgresRepository) GetByID(ctx context.Context, id int) (*entities.Usuario, error) {
	const q = `
		SELECT user_id, nombre, email, password, role_id, alias, telefono, residencia_id, eliminado, created_at, updated_at
		FROM usuario
		WHERE user_id = $1
	`
	var u entities.Usuario
	err := r.db.QueryRow(ctx, q, id).Scan(
		&u.ID, &u.Nombre, &u.Email, &u.PasswordHash, &u.RolID, &u.Alias, &u.Telefono, &u.ResidenciaID, &u.Eliminado, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UsuarioPostgresRepository) List(ctx context.Context) ([]entities.Usuario, error) {
	const q = `
		SELECT user_id, nombre, email, password, role_id, alias, telefono, residencia_id, eliminado, created_at, updated_at
		FROM usuario
		ORDER BY user_id DESC
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []entities.Usuario
	for rows.Next() {
		var u entities.Usuario
		if err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.PasswordHash, &u.RolID, &u.Alias, &u.Telefono, &u.ResidenciaID, &u.Eliminado, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}

func (r *UsuarioPostgresRepository) Delete(ctx context.Context, id int) error {
	const q = `DELETE FROM usuario WHERE user_id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *UsuarioPostgresRepository) FindByEmail(ctx context.Context, email string) (*entities.Usuario, error) {
	const q = `
		SELECT user_id, nombre, email, password, role_id, alias, telefono, residencia_id, eliminado, created_at, updated_at
		FROM usuario
		WHERE email = $1
	`
	var u entities.Usuario
	err := r.db.QueryRow(ctx, q, email).Scan(
		&u.ID, &u.Nombre, &u.Email, &u.PasswordHash, &u.RolID, &u.Alias, &u.Telefono, &u.ResidenciaID, &u.Eliminado, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
