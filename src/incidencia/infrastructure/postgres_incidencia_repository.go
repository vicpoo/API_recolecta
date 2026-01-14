// postgres_incidencia_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type PostgresIncidenciaRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresIncidenciaRepository() repositories.IIncidencia {
	pool := core.GetBD()
	return &PostgresIncidenciaRepository{pool: pool}
}

func (pg *PostgresIncidenciaRepository) Save(incidencia *entities.Incidencia) error {
	query := `
		INSERT INTO incidencia (
			punto_recoleccion_id, 
			conductor_id, 
			descripcion, 
			json_ruta, 
			fecha_reporte, 
			eliminado
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING incidencia_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		incidencia.PuntoRecoleccionID, 
		incidencia.ConductorID,
		incidencia.Descripcion,
		incidencia.JsonRuta,
		incidencia.FechaReporte,
		incidencia.Eliminado,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar la incidencia:", err)
		return err
	}
	
	incidencia.IncidenciaID = id
	return nil
}

func (pg *PostgresIncidenciaRepository) Update(incidencia *entities.Incidencia) error {
	query := `
		UPDATE incidencia
		SET 
			punto_recoleccion_id = $1, 
			conductor_id = $2, 
			descripcion = $3, 
			json_ruta = $4, 
			fecha_reporte = $5, 
			eliminado = $6
		WHERE incidencia_id = $7
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		incidencia.PuntoRecoleccionID, 
		incidencia.ConductorID,
		incidencia.Descripcion,
		incidencia.JsonRuta,
		incidencia.FechaReporte,
		incidencia.Eliminado,
		incidencia.IncidenciaID,
	)
	
	if err != nil {
		log.Println("Error al actualizar la incidencia:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("incidencia con ID %d no encontrada", incidencia.IncidenciaID)
	}

	return nil
}

func (pg *PostgresIncidenciaRepository) Delete(id int32) error {
	// Borrado físico (si lo necesitas)
	query := `
		DELETE FROM incidencia
		WHERE incidencia_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar la incidencia:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("incidencia con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresIncidenciaRepository) GetByID(id int32) (*entities.Incidencia, error) {
	query := `
		SELECT 
			incidencia_id,
			punto_recoleccion_id,
			conductor_id,
			descripcion,
			json_ruta,
			fecha_reporte,
			eliminado
		FROM incidencia
		WHERE incidencia_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var incidencia entities.Incidencia
	err := row.Scan(
		&incidencia.IncidenciaID,
		&incidencia.PuntoRecoleccionID,
		&incidencia.ConductorID,
		&incidencia.Descripcion,
		&incidencia.JsonRuta,
		&incidencia.FechaReporte,
		&incidencia.Eliminado,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("incidencia con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la incidencia por ID:", err)
		return nil, err
	}

	return &incidencia, nil
}

func (pg *PostgresIncidenciaRepository) GetAll() ([]entities.Incidencia, error) {
	query := `
		SELECT 
			incidencia_id,
			punto_recoleccion_id,
			conductor_id,
			descripcion,
			json_ruta,
			fecha_reporte,
			eliminado
		FROM incidencia
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todas las incidencias:", err)
		return nil, err
	}
	defer rows.Close()

	var incidencias []entities.Incidencia
	for rows.Next() {
		var incidencia entities.Incidencia
		err := rows.Scan(
			&incidencia.IncidenciaID,
			&incidencia.PuntoRecoleccionID,
			&incidencia.ConductorID,
			&incidencia.Descripcion,
			&incidencia.JsonRuta,
			&incidencia.FechaReporte,
			&incidencia.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear la incidencia:", err)
			return nil, err
		}
		incidencias = append(incidencias, incidencia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return incidencias, nil
}

func (pg *PostgresIncidenciaRepository) GetByConductorID(conductorID int32) ([]entities.Incidencia, error) {
	query := `
		SELECT 
			incidencia_id,
			punto_recoleccion_id,
			conductor_id,
			descripcion,
			json_ruta,
			fecha_reporte,
			eliminado
		FROM incidencia
		WHERE conductor_id = $1
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, conductorID)
	if err != nil {
		log.Println("Error al obtener incidencias por conductor ID:", err)
		return nil, err
	}
	defer rows.Close()

	var incidencias []entities.Incidencia
	for rows.Next() {
		var incidencia entities.Incidencia
		err := rows.Scan(
			&incidencia.IncidenciaID,
			&incidencia.PuntoRecoleccionID,
			&incidencia.ConductorID,
			&incidencia.Descripcion,
			&incidencia.JsonRuta,
			&incidencia.FechaReporte,
			&incidencia.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear la incidencia:", err)
			return nil, err
		}
		incidencias = append(incidencias, incidencia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return incidencias, nil
}

func (pg *PostgresIncidenciaRepository) GetByPuntoRecoleccionID(puntoRecoleccionID int32) ([]entities.Incidencia, error) {
	query := `
		SELECT 
			incidencia_id,
			punto_recoleccion_id,
			conductor_id,
			descripcion,
			json_ruta,
			fecha_reporte,
			eliminado
		FROM incidencia
		WHERE punto_recoleccion_id = $1
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, puntoRecoleccionID)
	if err != nil {
		log.Println("Error al obtener incidencias por punto de recolección ID:", err)
		return nil, err
	}
	defer rows.Close()

	var incidencias []entities.Incidencia
	for rows.Next() {
		var incidencia entities.Incidencia
		err := rows.Scan(
			&incidencia.IncidenciaID,
			&incidencia.PuntoRecoleccionID,
			&incidencia.ConductorID,
			&incidencia.Descripcion,
			&incidencia.JsonRuta,
			&incidencia.FechaReporte,
			&incidencia.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear la incidencia:", err)
			return nil, err
		}
		incidencias = append(incidencias, incidencia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return incidencias, nil
}

func (pg *PostgresIncidenciaRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Incidencia, error) {
	// Parsear fechas
	startTime, err := time.Parse("2006-01-02", fechaInicio)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha_inicio inválido: %v", err)
	}
	
	endTime, err := time.Parse("2006-01-02", fechaFin)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha_fin inválido: %v", err)
	}
	
	// Ajustar para incluir todo el día final
	endTime = endTime.Add(24 * time.Hour)
	
	query := `
		SELECT 
			incidencia_id,
			punto_recoleccion_id,
			conductor_id,
			descripcion,
			json_ruta,
			fecha_reporte,
			eliminado
		FROM incidencia
		WHERE fecha_reporte BETWEEN $1 AND $2
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, startTime, endTime)
	if err != nil {
		log.Println("Error al obtener incidencias por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var incidencias []entities.Incidencia
	for rows.Next() {
		var incidencia entities.Incidencia
		err := rows.Scan(
			&incidencia.IncidenciaID,
			&incidencia.PuntoRecoleccionID,
			&incidencia.ConductorID,
			&incidencia.Descripcion,
			&incidencia.JsonRuta,
			&incidencia.FechaReporte,
			&incidencia.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear la incidencia:", err)
			return nil, err
		}
		incidencias = append(incidencias, incidencia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return incidencias, nil
}