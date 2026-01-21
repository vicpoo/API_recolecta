package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/rol/domain"
)

type RolRepository struct {
	db *pgxpool.Pool
}

func NewRolRepository(db *pgxpool.Pool) *RolRepository {
	return &RolRepository{db}
}

func (r *RolRepository) Create(nombre string) error {
	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO rol (nombre) VALUES ($1)`,
		nombre,
	)
	return err
}

func (r *RolRepository) List() ([]domain.Rol, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT role_id, nombre, eliminado FROM rol WHERE eliminado = false`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Rol
	for rows.Next() {
		var rol domain.Rol
		if err := rows.Scan(&rol.ID, &rol.Nombre, &rol.Eliminado); err != nil {
			return nil, err
		}
		roles = append(roles, rol)
	}
	return roles, nil
}

func (r *RolRepository) Update(id int, nombre string) error {
	_, err := r.db.Exec(
		context.Background(),
		`UPDATE rol SET nombre=$1 WHERE role_id=$2`,
		nombre, id,
	)
	return err
}
