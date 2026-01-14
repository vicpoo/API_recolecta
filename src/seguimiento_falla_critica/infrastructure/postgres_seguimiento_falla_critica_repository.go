// postgres_seguimiento_falla_critica_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/domain/entities"
)

type PostgresSeguimientoFallaCriticaRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresSeguimientoFallaCriticaRepository() repositories.ISeguimientoFallaCritica {
	pool := core.GetBD()
	return &PostgresSeguimientoFallaCriticaRepository{pool: pool}
}

func (pg *PostgresSeguimientoFallaCriticaRepository) Save(seguimiento *entities.SeguimientoFallaCritica) error {
	query := `
		INSERT INTO seguimiento_falla_critica (falla_id, comentario, created_at)
		VALUES ($1, $2, $3)
		RETURNING seguimiento_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		seguimiento.FallaID, 
		seguimiento.Comentario,
		seguimiento.CreatedAt,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar el seguimiento de falla crítica:", err)
		return err
	}
	
	seguimiento.SeguimientoID = id
	return nil
}

func (pg *PostgresSeguimientoFallaCriticaRepository) Update(seguimiento *entities.SeguimientoFallaCritica) error {
	query := `
		UPDATE seguimiento_falla_critica
		SET falla_id = $1, comentario = $2
		WHERE seguimiento_id = $3
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		seguimiento.FallaID, 
		seguimiento.Comentario,
		seguimiento.SeguimientoID,
	)
	
	if err != nil {
		log.Println("Error al actualizar el seguimiento de falla crítica:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("seguimiento de falla crítica con ID %d no encontrado", seguimiento.SeguimientoID)
	}

	return nil
}

func (pg *PostgresSeguimientoFallaCriticaRepository) Delete(id int32) error {
	query := `
		DELETE FROM seguimiento_falla_critica
		WHERE seguimiento_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar el seguimiento de falla crítica:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("seguimiento de falla crítica con ID %d no encontrado", id)
	}

	return nil
}

func (pg *PostgresSeguimientoFallaCriticaRepository) GetByID(id int32) (*entities.SeguimientoFallaCritica, error) {
	query := `
		SELECT seguimiento_id, falla_id, comentario, created_at
		FROM seguimiento_falla_critica
		WHERE seguimiento_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var seguimiento entities.SeguimientoFallaCritica
	err := row.Scan(
		&seguimiento.SeguimientoID,
		&seguimiento.FallaID,
		&seguimiento.Comentario,
		&seguimiento.CreatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("seguimiento de falla crítica con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el seguimiento de falla crítica por ID:", err)
		return nil, err
	}

	return &seguimiento, nil
}

func (pg *PostgresSeguimientoFallaCriticaRepository) GetAll() ([]entities.SeguimientoFallaCritica, error) {
	query := `
		SELECT seguimiento_id, falla_id, comentario, created_at
		FROM seguimiento_falla_critica
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todos los seguimientos de falla crítica:", err)
		return nil, err
	}
	defer rows.Close()

	var seguimientos []entities.SeguimientoFallaCritica
	for rows.Next() {
		var seguimiento entities.SeguimientoFallaCritica
		err := rows.Scan(
			&seguimiento.SeguimientoID,
			&seguimiento.FallaID,
			&seguimiento.Comentario,
			&seguimiento.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el seguimiento de falla crítica:", err)
			return nil, err
		}
		seguimientos = append(seguimientos, seguimiento)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return seguimientos, nil
}

func (pg *PostgresSeguimientoFallaCriticaRepository) GetByFallaID(fallaID int32) ([]entities.SeguimientoFallaCritica, error) {
	query := `
		SELECT seguimiento_id, falla_id, comentario, created_at
		FROM seguimiento_falla_critica
		WHERE falla_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fallaID)
	if err != nil {
		log.Println("Error al obtener seguimientos de falla crítica por falla ID:", err)
		return nil, err
	}
	defer rows.Close()

	var seguimientos []entities.SeguimientoFallaCritica
	for rows.Next() {
		var seguimiento entities.SeguimientoFallaCritica
		err := rows.Scan(
			&seguimiento.SeguimientoID,
			&seguimiento.FallaID,
			&seguimiento.Comentario,
			&seguimiento.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el seguimiento de falla crítica:", err)
			return nil, err
		}
		seguimientos = append(seguimientos, seguimiento)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return seguimientos, nil
}

func (pg *PostgresSeguimientoFallaCriticaRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.SeguimientoFallaCritica, error) {
	query := `
		SELECT seguimiento_id, falla_id, comentario, created_at
		FROM seguimiento_falla_critica
		WHERE created_at >= $1 AND created_at <= $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener seguimientos de falla crítica por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var seguimientos []entities.SeguimientoFallaCritica
	for rows.Next() {
		var seguimiento entities.SeguimientoFallaCritica
		err := rows.Scan(
			&seguimiento.SeguimientoID,
			&seguimiento.FallaID,
			&seguimiento.Comentario,
			&seguimiento.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el seguimiento de falla crítica:", err)
			return nil, err
		}
		seguimientos = append(seguimientos, seguimiento)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return seguimientos, nil
}