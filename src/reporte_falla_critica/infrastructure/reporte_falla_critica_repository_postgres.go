// postgres_reporte_falla_critica_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type PostgresReporteFallaCriticaRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresReporteFallaCriticaRepository() repositories.IReporteFallaCritica {
	pool := core.GetBD()
	return &PostgresReporteFallaCriticaRepository{pool: pool}
}

func (pg *PostgresReporteFallaCriticaRepository) Save(reporteFallaCritica *entities.ReporteFallaCritica) error {
	query := `
		INSERT INTO reporte_falla_critica (camion_id, conductor_id, descripcion, created_at, eliminado)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING falla_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		reporteFallaCritica.CamionID, 
		reporteFallaCritica.ConductorID,
		reporteFallaCritica.Descripcion,
		reporteFallaCritica.CreatedAt,
		reporteFallaCritica.Eliminado,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar el reporte de falla crítica:", err)
		return err
	}
	
	reporteFallaCritica.FallaID = id
	return nil
}

func (pg *PostgresReporteFallaCriticaRepository) Update(reporteFallaCritica *entities.ReporteFallaCritica) error {
	query := `
		UPDATE reporte_falla_critica
		SET camion_id = $1, conductor_id = $2, descripcion = $3, eliminado = $4
		WHERE falla_id = $5
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		reporteFallaCritica.CamionID, 
		reporteFallaCritica.ConductorID,
		reporteFallaCritica.Descripcion,
		reporteFallaCritica.Eliminado,
		reporteFallaCritica.FallaID,
	)
	
	if err != nil {
		log.Println("Error al actualizar el reporte de falla crítica:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("reporte de falla crítica con ID %d no encontrado", reporteFallaCritica.FallaID)
	}

	return nil
}

func (pg *PostgresReporteFallaCriticaRepository) Delete(id int32) error {
	// Borrado lógico (soft delete)
	query := `
		UPDATE reporte_falla_critica
		SET eliminado = true
		WHERE falla_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar (marcar como eliminado) el reporte de falla crítica:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("reporte de falla crítica con ID %d no encontrado", id)
	}

	return nil
}

func (pg *PostgresReporteFallaCriticaRepository) GetByID(id int32) (*entities.ReporteFallaCritica, error) {
	query := `
		SELECT falla_id, camion_id, conductor_id, descripcion, created_at, eliminado
		FROM reporte_falla_critica
		WHERE falla_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var reporte entities.ReporteFallaCritica
	err := row.Scan(
		&reporte.FallaID,
		&reporte.CamionID,
		&reporte.ConductorID,
		&reporte.Descripcion,
		&reporte.CreatedAt,
		&reporte.Eliminado,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("reporte de falla crítica con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el reporte de falla crítica por ID:", err)
		return nil, err
	}

	return &reporte, nil
}

func (pg *PostgresReporteFallaCriticaRepository) GetAll() ([]entities.ReporteFallaCritica, error) {
	query := `
		SELECT falla_id, camion_id, conductor_id, descripcion, created_at, eliminado
		FROM reporte_falla_critica
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todos los reportes de falla crítica:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteFallaCritica
	for rows.Next() {
		var reporte entities.ReporteFallaCritica
		err := rows.Scan(
			&reporte.FallaID,
			&reporte.CamionID,
			&reporte.ConductorID,
			&reporte.Descripcion,
			&reporte.CreatedAt,
			&reporte.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de falla crítica:", err)
			return nil, err
		}
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteFallaCriticaRepository) GetByCamionID(camionID int32) ([]entities.ReporteFallaCritica, error) {
	query := `
		SELECT falla_id, camion_id, conductor_id, descripcion, created_at, eliminado
		FROM reporte_falla_critica
		WHERE camion_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, camionID)
	if err != nil {
		log.Println("Error al obtener reportes de falla crítica por camión ID:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteFallaCritica
	for rows.Next() {
		var reporte entities.ReporteFallaCritica
		err := rows.Scan(
			&reporte.FallaID,
			&reporte.CamionID,
			&reporte.ConductorID,
			&reporte.Descripcion,
			&reporte.CreatedAt,
			&reporte.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de falla crítica:", err)
			return nil, err
		}
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteFallaCriticaRepository) GetByConductorID(conductorID int32) ([]entities.ReporteFallaCritica, error) {
	query := `
		SELECT falla_id, camion_id, conductor_id, descripcion, created_at, eliminado
		FROM reporte_falla_critica
		WHERE conductor_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, conductorID)
	if err != nil {
		log.Println("Error al obtener reportes de falla crítica por conductor ID:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteFallaCritica
	for rows.Next() {
		var reporte entities.ReporteFallaCritica
		err := rows.Scan(
			&reporte.FallaID,
			&reporte.CamionID,
			&reporte.ConductorID,
			&reporte.Descripcion,
			&reporte.CreatedAt,
			&reporte.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de falla crítica:", err)
			return nil, err
		}
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteFallaCriticaRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.ReporteFallaCritica, error) {
	query := `
		SELECT falla_id, camion_id, conductor_id, descripcion, created_at, eliminado
		FROM reporte_falla_critica
		WHERE created_at >= $1 AND created_at <= $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener reportes de falla crítica por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteFallaCritica
	for rows.Next() {
		var reporte entities.ReporteFallaCritica
		err := rows.Scan(
			&reporte.FallaID,
			&reporte.CamionID,
			&reporte.ConductorID,
			&reporte.Descripcion,
			&reporte.CreatedAt,
			&reporte.Eliminado,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de falla crítica:", err)
			return nil, err
		}
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}