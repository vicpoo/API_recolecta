package adapters

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/Ruta/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres() *Postgres {
	conn, _ := core.ConnectPostgres()
	return &Postgres{conn: conn}
}

func (pg *Postgres) Save(ruta *entities.Ruta) error {
	ruta.CreatedAt = time.Now()
	sql := `
		INSERT INTO ruta (nombre, descripcion, json_ruta, eliminado, created_at)
		VALUES ($1, $2, $3, false, $4)
		RETURNING ruta_id
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		ruta.Nombre,
		ruta.Descripcion,
		ruta.JsonRuta, // Ya es string, se guarda directamente
		ruta.CreatedAt,
	).Scan(&ruta.RutaID)

	if err != nil {
		return fmt.Errorf("error al guardar ruta: %w", err)
	}
	return nil
}

func (pg *Postgres) ListAll() ([]entities.Ruta, error) {
	sql := `
		SELECT ruta_id, nombre, descripcion, json_ruta, eliminado, created_at
		FROM ruta
		ORDER BY ruta_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error al consultar rutas: %w", err)
	}
	defer rows.Close()

	var rutas []entities.Ruta
	for rows.Next() {
		var r entities.Ruta

		err := rows.Scan(
			&r.RutaID,
			&r.Nombre,
			&r.Descripcion,
			&r.JsonRuta, // Se lee directamente como string
			&r.Eliminado,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear ruta: %w", err)
		}

		rutas = append(rutas, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración: %w", err)
	}

	return rutas, nil
}

func (pg *Postgres) GetById(id int32) (*entities.Ruta, error) {
	sql := `
		SELECT ruta_id, nombre, descripcion, json_ruta, eliminado, created_at
		FROM ruta
		WHERE ruta_id = $1
	`

	var r entities.Ruta

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(
		&r.RutaID,
		&r.Nombre,
		&r.Descripcion,
		&r.JsonRuta, // Se lee directamente como string
		&r.Eliminado,
		&r.CreatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, errors.New("ruta no encontrada")
	}
	if err != nil {
		return nil, fmt.Errorf("error al obtener ruta: %w", err)
	}

	return &r, nil
}

func (pg *Postgres) Update(ruta *entities.Ruta) error {
	sql := `
		UPDATE ruta
		SET nombre = $1, descripcion = $2, json_ruta = $3
		WHERE ruta_id = $4 AND eliminado = false
	`

	cmd, err := pg.conn.Exec(
		context.Background(),
		sql,
		ruta.Nombre,
		ruta.Descripcion,
		ruta.JsonRuta, // Ya es string, se guarda directamente
		ruta.RutaID,
	)
	if err != nil {
		return fmt.Errorf("error al actualizar ruta: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("ruta no encontrada o ya eliminada")
	}

	return nil
}

func (pg *Postgres) Delete(id int32) error {
	sql := `
		UPDATE ruta
		SET eliminado = true
		WHERE ruta_id = $1 AND eliminado = false
	`

	cmd, err := pg.conn.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("error al eliminar ruta: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("ruta no encontrada o ya eliminada")
	}

	return nil
}

func (pg *Postgres) GetActivas() ([]entities.Ruta, error) {
	sql := `
		SELECT ruta_id, nombre, descripcion, json_ruta, eliminado, created_at
		FROM ruta
		WHERE eliminado = false
		ORDER BY ruta_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error al consultar rutas activas: %w", err)
	}
	defer rows.Close()

	var rutas []entities.Ruta
	for rows.Next() {
		var r entities.Ruta

		err := rows.Scan(
			&r.RutaID,
			&r.Nombre,
			&r.Descripcion,
			&r.JsonRuta, // Se lee directamente como string
			&r.Eliminado,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear ruta: %w", err)
		}

		rutas = append(rutas, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración: %w", err)
	}

	return rutas, nil
}