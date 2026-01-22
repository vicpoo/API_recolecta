package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/rol/domain/entities"
)

type RolRepository struct {
	db *pgxpool.Pool
}

func NewRolRepository(db *pgxpool.Pool) *RolRepository {
	return &RolRepository{db}
}

// Create recibe solo el nombre (string)
func (r *RolRepository) Create(nombre string) error {
	query := `INSERT INTO rol (nombre) VALUES ($1)`

	_, err := r.db.Exec(
		context.Background(),
		query,
		nombre,
	)

	return err
}

// List en lugar de GetAll
func (r *RolRepository) List() ([]entities.Rol, error) {
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

	var roles []entities.Rol

	for rows.Next() {
		var role entities.Rol
		if err := rows.Scan(
			&role.ID,
			&role.Nombre,
			&role.Eliminado,
		); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

// Update recibe id e nombre (no el objeto completo)
func (r *RolRepository) Update(id int, nombre string) error {
	query := `
		UPDATE rol
		SET nombre = $1
		WHERE role_id = $2 AND eliminado = false
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		nombre,
		id,
	)

	return err
}

// Delete marca el rol como eliminado (soft delete)
func (r *RolRepository) Delete(id int) error {
	query := `
		UPDATE rol
		SET eliminado = true
		WHERE role_id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}

// Métodos auxiliares (opcionales, no están en la interfaz)
func (r *RolRepository) GetByID(id int) (*entities.Rol, error) {
	query := `
		SELECT role_id, nombre, eliminado
		FROM rol
		WHERE role_id = $1 AND eliminado = false
	`

	row := r.db.QueryRow(context.Background(), query, id)

	var role entities.Rol
	err := row.Scan(
		&role.ID,
		&role.Nombre,
		&role.Eliminado,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}