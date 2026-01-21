package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/rol/domain"
)

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) Create(role *domain.Role) error {
	query := `INSERT INTO rol (nombre) VALUES ($1)`

	_, err := r.db.Exec(
		context.Background(),
		query,
		role.Nombre,
	)

	return err
}

func (r *RoleRepository) GetByID(id int) (*domain.Role, error) {
	query := `
		SELECT role_id, nombre, eliminado
		FROM rol
		WHERE role_id = $1 AND eliminado = false
	`

	row := r.db.QueryRow(context.Background(), query, id)

	var role domain.Role
	err := row.Scan(
		&role.RoleID,
		&role.Nombre,
		&role.Eliminado,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) GetAll() ([]domain.Role, error) {
	query := `
		SELECT role_id, nombre, eliminado
		FROM rol
		WHERE eliminado = false
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role

	for rows.Next() {
		var role domain.Role
		if err := rows.Scan(
			&role.RoleID,
			&role.Nombre,
			&role.Eliminado,
		); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RoleRepository) Update(role *domain.Role) error {
	query := `
		UPDATE rol
		SET nombre = $1
		WHERE role_id = $2 AND eliminado = false
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		role.Nombre,
		role.RoleID,
	)

	return err
}
