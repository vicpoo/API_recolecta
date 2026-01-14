// postgres_reporte_conductor_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/reporte_conductor/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type PostgresReporteConductorRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresReporteConductorRepository() repositories.IReporteConductor {
	pool := core.GetBD()
	return &PostgresReporteConductorRepository{pool: pool}
}

func (pg *PostgresReporteConductorRepository) Save(reporteConductor *entities.ReporteConductor) error {
	query := `
		INSERT INTO reporte_conductor (conductor_id, camion_id, ruta_id, descripcion, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING reporte_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		reporteConductor.ConductorID, 
		reporteConductor.CamionID,
		reporteConductor.RutaID,
		reporteConductor.Descripcion,
		reporteConductor.CreatedAt,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar el reporte del conductor:", err)
		return err
	}
	
	reporteConductor.SetReporteID(id)
	return nil
}

func (pg *PostgresReporteConductorRepository) Update(reporteConductor *entities.ReporteConductor) error {
	query := `
		UPDATE reporte_conductor
		SET conductor_id = $1, camion_id = $2, ruta_id = $3, descripcion = $4
		WHERE reporte_id = $5
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		reporteConductor.ConductorID, 
		reporteConductor.CamionID,
		reporteConductor.RutaID,
		reporteConductor.Descripcion,
		reporteConductor.ReporteID,
	)
	
	if err != nil {
		log.Println("Error al actualizar el reporte del conductor:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("reporte del conductor con ID %d no encontrado", reporteConductor.ReporteID)
	}

	return nil
}

func (pg *PostgresReporteConductorRepository) Delete(id int32) error {
	query := `
		DELETE FROM reporte_conductor
		WHERE reporte_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar el reporte del conductor:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("reporte del conductor con ID %d no encontrado", id)
	}

	return nil
}

func (pg *PostgresReporteConductorRepository) GetByID(id int32) (*entities.ReporteConductor, error) {
	query := `
		SELECT reporte_id, conductor_id, camion_id, ruta_id, descripcion, created_at
		FROM reporte_conductor
		WHERE reporte_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var reporteConductor entities.ReporteConductor
	var createdAt time.Time
	err := row.Scan(
		&reporteConductor.ReporteID,
		&reporteConductor.ConductorID,
		&reporteConductor.CamionID,
		&reporteConductor.RutaID,
		&reporteConductor.Descripcion,
		&createdAt,
	)
	
	reporteConductor.SetCreatedAt(createdAt)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("reporte del conductor con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el reporte del conductor por ID:", err)
		return nil, err
	}

	return &reporteConductor, nil
}

func (pg *PostgresReporteConductorRepository) GetAll() ([]entities.ReporteConductor, error) {
	query := `
		SELECT reporte_id, conductor_id, camion_id, ruta_id, descripcion, created_at
		FROM reporte_conductor
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todos los reportes de conductores:", err)
		return nil, err
	}
	defer rows.Close()

	var reportesConductor []entities.ReporteConductor
	for rows.Next() {
		var reporte entities.ReporteConductor
		var createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.ConductorID,
			&reporte.CamionID,
			&reporte.RutaID,
			&reporte.Descripcion,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte del conductor:", err)
			return nil, err
		}
		reporte.SetCreatedAt(createdAt)
		reportesConductor = append(reportesConductor, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportesConductor, nil
}

func (pg *PostgresReporteConductorRepository) GetByCamionID(camionID int32) ([]entities.ReporteConductor, error) {
	query := `
		SELECT reporte_id, conductor_id, camion_id, ruta_id, descripcion, created_at
		FROM reporte_conductor
		WHERE camion_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, camionID)
	if err != nil {
		log.Println("Error al obtener reportes por camión ID:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteConductor
	for rows.Next() {
		var reporte entities.ReporteConductor
		var createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.ConductorID,
			&reporte.CamionID,
			&reporte.RutaID,
			&reporte.Descripcion,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte del conductor:", err)
			return nil, err
		}
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteConductorRepository) GetByConductorID(conductorID int32) ([]entities.ReporteConductor, error) {
	query := `
		SELECT reporte_id, conductor_id, camion_id, ruta_id, descripcion, created_at
		FROM reporte_conductor
		WHERE conductor_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, conductorID)
	if err != nil {
		log.Println("Error al obtener reportes por conductor ID:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteConductor
	for rows.Next() {
		var reporte entities.ReporteConductor
		var createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.ConductorID,
			&reporte.CamionID,
			&reporte.RutaID,
			&reporte.Descripcion,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte del conductor:", err)
			return nil, err
		}
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteConductorRepository) GetByRutaID(rutaID int32) ([]entities.ReporteConductor, error) {
	query := `
		SELECT reporte_id, conductor_id, camion_id, ruta_id, descripcion, created_at
		FROM reporte_conductor
		WHERE ruta_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, rutaID)
	if err != nil {
		log.Println("Error al obtener reportes por ruta ID:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteConductor
	for rows.Next() {
		var reporte entities.ReporteConductor
		var createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.ConductorID,
			&reporte.CamionID,
			&reporte.RutaID,
			&reporte.Descripcion,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte del conductor:", err)
			return nil, err
		}
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteConductorRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.ReporteConductor, error) {
	query := `
		SELECT reporte_id, conductor_id, camion_id, ruta_id, descripcion, created_at
		FROM reporte_conductor
		WHERE created_at >= $1 AND created_at <= $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener reportes por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteConductor
	for rows.Next() {
		var reporte entities.ReporteConductor
		var createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.ConductorID,
			&reporte.CamionID,
			&reporte.RutaID,
			&reporte.Descripcion,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte del conductor:", err)
			return nil, err
		}
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}