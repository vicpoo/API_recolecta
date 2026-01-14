// postgres_tipo_mantenimiento_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/domain/entities"
)

type PostgresTipoMantenimientoRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresTipoMantenimientoRepository() repositories.ITipoMantenimiento {
	pool := core.GetBD()
	return &PostgresTipoMantenimientoRepository{pool: pool}
}

func (pg *PostgresTipoMantenimientoRepository) Save(tipoMantenimiento *entities.TipoMantenimiento) error {
	query := `
		INSERT INTO tipo_mantenimiento (nombre, categoria, eliminado)
		VALUES ($1, $2, $3)
		RETURNING tipo_mantenimiento_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		tipoMantenimiento.Nombre, 
		tipoMantenimiento.Categoria,
		tipoMantenimiento.Eliminado,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar el tipo de mantenimiento:", err)
		return err
	}
	
	tipoMantenimiento.ID = id
	return nil
}

func (pg *PostgresTipoMantenimientoRepository) Update(tipoMantenimiento *entities.TipoMantenimiento) error {
	query := `
		UPDATE tipo_mantenimiento
		SET nombre = $1, categoria = $2, eliminado = $3
		WHERE tipo_mantenimiento_id = $4
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		tipoMantenimiento.Nombre, 
		tipoMantenimiento.Categoria,
		tipoMantenimiento.Eliminado,
		tipoMantenimiento.ID,
	)
	
	if err != nil {
		log.Println("Error al actualizar el tipo de mantenimiento:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("tipo de mantenimiento con ID %d no encontrado", tipoMantenimiento.ID)
	}

	return nil
}

func (pg *PostgresTipoMantenimientoRepository) Delete(id int32) error {
	// Borrado lógico
	query := `
		UPDATE tipo_mantenimiento
		SET eliminado = true
		WHERE tipo_mantenimiento_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar (marcar como eliminado) el tipo de mantenimiento:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("tipo de mantenimiento con ID %d no encontrado", id)
	}

	return nil
}

func (pg *PostgresTipoMantenimientoRepository) GetByID(id int32) (*entities.TipoMantenimiento, error) {
	query := `
		SELECT tipo_mantenimiento_id, nombre, categoria, eliminado
		FROM tipo_mantenimiento
		WHERE tipo_mantenimiento_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var tipoMantenimiento entities.TipoMantenimiento
	err := row.Scan(
		&tipoMantenimiento.ID,
		&tipoMantenimiento.Nombre,
		&tipoMantenimiento.Categoria,
		&tipoMantenimiento.Eliminado,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("tipo de mantenimiento con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el tipo de mantenimiento por ID:", err)
		return nil, err
	}

	return &tipoMantenimiento, nil
}

func (pg *PostgresTipoMantenimientoRepository) GetAll() ([]entities.TipoMantenimiento, error) {
	query := `
		SELECT tipo_mantenimiento_id, nombre, categoria, eliminado
		FROM tipo_mantenimiento
		ORDER BY tipo_mantenimiento_id
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todos los tipos de mantenimiento:", err)
		return nil, err
	}
	defer rows.Close()

	var tiposMantenimiento []entities.TipoMantenimiento
	for rows.Next() {
		var tipo entities.TipoMantenimiento
		err := rows.Scan(
			&tipo.ID,
			&tipo.Nombre,
			&tipo.Categoria,
			&tipo.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear el tipo de mantenimiento:", err)
			return nil, err
		}
		tiposMantenimiento = append(tiposMantenimiento, tipo)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return tiposMantenimiento, nil
}