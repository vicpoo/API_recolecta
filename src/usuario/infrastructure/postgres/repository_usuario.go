package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/usuario/domain"
)

type UsuarioRepository struct {
	db *pgxpool.Pool
}

func NewUsuarioRepository(db *pgxpool.Pool) *UsuarioRepository {
	return &UsuarioRepository{db}
}

func (r *UsuarioRepository) Create(u *domain.Usuario) error {
	query := `
		INSERT INTO usuario
		(nombre, alias, telefono, email, password, role_id, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		u.Nombre,
		u.Alias,
		u.Telefono,
		u.Email,
		u.Password,
		u.RoleID,
		u.CreatedAt,
		u.UpdatedAt,
	)

	return err
}

func (r *UsuarioRepository) GetByEmail(email string) (*domain.Usuario, error) {
	query := `
		SELECT user_id, nombre, alias, telefono, email, password,
		role_id, eliminado, created_at, updated_at
		FROM usuario
		WHERE email = $1 AND eliminado = false
	`

	row := r.db.QueryRow(context.Background(), query, email)

	var u domain.Usuario
	err := row.Scan(
		&u.UserID,
		&u.Nombre,
		&u.Alias,
		&u.Telefono,
		&u.Email,
		&u.Password,
		&u.RoleID,
		&u.Eliminado,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}


func (r *UsuarioRepository) GetByID(id int) (*domain.Usuario, error) {
	query := `
		SELECT user_id, nombre, alias, telefono, email, role_id,
		    eliminado, created_at, updated_at
		FROM usuario
		WHERE user_id = $1 AND eliminado = false
	`

	row := r.db.QueryRow(context.Background(), query, id)

	var u domain.Usuario
	err := row.Scan(
		&u.UserID,
		&u.Nombre,
		&u.Alias,
		&u.Telefono,
		&u.Email,
		&u.RoleID,
		&u.Eliminado,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UsuarioRepository) GetAll() ([]domain.Usuario, error) {
	query := `
		SELECT user_id, nombre, alias, telefono, email, role_id,
		eliminado, created_at, updated_at
		FROM usuario
		WHERE eliminado = false
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []domain.Usuario

	for rows.Next() {
		var u domain.Usuario
		if err := rows.Scan(
			&u.UserID,
			&u.Nombre,
			&u.Alias,
			&u.Telefono,
			&u.Email,
			&u.RoleID,
			&u.Eliminado,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}

func (r *UsuarioRepository) Delete(id int) error {
	query := `
		UPDATE usuario
		SET eliminado = true, updated_at = NOW()
		WHERE user_id = $1
	`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}