// postgres_reporte_mantenimiento_generado_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type PostgresReporteMantenimientoGeneradoRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresReporteMantenimientoGeneradoRepository() repositories.IReporteMantenimientoGenerado {
	pool := core.GetBD()
	return &PostgresReporteMantenimientoGeneradoRepository{pool: pool}
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) Save(reporte *entities.ReporteMantenimientoGenerado) error {
	query := `
		INSERT INTO reporte_mantenimiento_generado (coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING reporte_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		reporte.CoordinadorID, 
		reporte.FechaDesde,
		reporte.FechaHasta,
		reporte.Observaciones,
		reporte.CreatedAt,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar el reporte de mantenimiento generado:", err)
		return err
	}
	
	reporte.SetReporteID(id)
	return nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) Update(reporte *entities.ReporteMantenimientoGenerado) error {
	query := `
		UPDATE reporte_mantenimiento_generado
		SET coordinador_id = $1, fecha_desde = $2, fecha_hasta = $3, observaciones = $4
		WHERE reporte_id = $5
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		reporte.CoordinadorID, 
		reporte.FechaDesde,
		reporte.FechaHasta,
		reporte.Observaciones,
		reporte.ReporteID,
	)
	
	if err != nil {
		log.Println("Error al actualizar el reporte de mantenimiento generado:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("reporte de mantenimiento generado con ID %d no encontrado", reporte.ReporteID)
	}

	return nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) Delete(id int32) error {
	query := `
		DELETE FROM reporte_mantenimiento_generado
		WHERE reporte_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar el reporte de mantenimiento generado:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("reporte de mantenimiento generado con ID %d no encontrado", id)
	}

	return nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) GetByID(id int32) (*entities.ReporteMantenimientoGenerado, error) {
	query := `
		SELECT reporte_id, coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at
		FROM reporte_mantenimiento_generado
		WHERE reporte_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var reporte entities.ReporteMantenimientoGenerado
	var fechaDesde, fechaHasta, createdAt time.Time
	
	err := row.Scan(
		&reporte.ReporteID,
		&reporte.CoordinadorID,
		&fechaDesde,
		&fechaHasta,
		&reporte.Observaciones,
		&createdAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("reporte de mantenimiento generado con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el reporte de mantenimiento generado por ID:", err)
		return nil, err
	}

	reporte.SetFechaDesde(fechaDesde)
	reporte.SetFechaHasta(fechaHasta)
	reporte.SetCreatedAt(createdAt)
	
	return &reporte, nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) GetAll() ([]entities.ReporteMantenimientoGenerado, error) {
	query := `
		SELECT reporte_id, coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at
		FROM reporte_mantenimiento_generado
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todos los reportes de mantenimiento generados:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteMantenimientoGenerado
	for rows.Next() {
		var reporte entities.ReporteMantenimientoGenerado
		var fechaDesde, fechaHasta, createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.CoordinadorID,
			&fechaDesde,
			&fechaHasta,
			&reporte.Observaciones,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de mantenimiento generado:", err)
			return nil, err
		}
		reporte.SetFechaDesde(fechaDesde)
		reporte.SetFechaHasta(fechaHasta)
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) GetByCoordinadorID(coordinadorID int32) ([]entities.ReporteMantenimientoGenerado, error) {
	query := `
		SELECT reporte_id, coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at
		FROM reporte_mantenimiento_generado
		WHERE coordinador_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, coordinadorID)
	if err != nil {
		log.Println("Error al obtener reportes por coordinador ID:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteMantenimientoGenerado
	for rows.Next() {
		var reporte entities.ReporteMantenimientoGenerado
		var fechaDesde, fechaHasta, createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.CoordinadorID,
			&fechaDesde,
			&fechaHasta,
			&reporte.Observaciones,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de mantenimiento generado:", err)
			return nil, err
		}
		reporte.SetFechaDesde(fechaDesde)
		reporte.SetFechaHasta(fechaHasta)
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error) {
	query := `
		SELECT reporte_id, coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at
		FROM reporte_mantenimiento_generado
		WHERE fecha_desde >= $1 AND fecha_hasta <= $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener reportes por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteMantenimientoGenerado
	for rows.Next() {
		var reporte entities.ReporteMantenimientoGenerado
		var fechaDesde, fechaHasta, createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.CoordinadorID,
			&fechaDesde,
			&fechaHasta,
			&reporte.Observaciones,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de mantenimiento generado:", err)
			return nil, err
		}
		reporte.SetFechaDesde(fechaDesde)
		reporte.SetFechaHasta(fechaHasta)
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

func (pg *PostgresReporteMantenimientoGeneradoRepository) GetByFechaGeneracionRange(fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error) {
	query := `
		SELECT reporte_id, coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at
		FROM reporte_mantenimiento_generado
		WHERE created_at >= $1 AND created_at <= $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener reportes por rango de fechas de generación:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteMantenimientoGenerado
	for rows.Next() {
		var reporte entities.ReporteMantenimientoGenerado
		var fechaDesde, fechaHasta, createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.CoordinadorID,
			&fechaDesde,
			&fechaHasta,
			&reporte.Observaciones,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de mantenimiento generado:", err)
			return nil, err
		}
		reporte.SetFechaDesde(fechaDesde)
		reporte.SetFechaHasta(fechaHasta)
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}

// Método adicional (si decides implementar GetByCoordinadorYFecha)
func (pg *PostgresReporteMantenimientoGeneradoRepository) GetByCoordinadorYFecha(coordinadorID int32, fechaInicio, fechaFin string) ([]entities.ReporteMantenimientoGenerado, error) {
	query := `
		SELECT reporte_id, coordinador_id, fecha_desde, fecha_hasta, observaciones, created_at
		FROM reporte_mantenimiento_generado
		WHERE coordinador_id = $1 AND fecha_desde >= $2 AND fecha_hasta <= $3
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, coordinadorID, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener reportes por coordinador y rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var reportes []entities.ReporteMantenimientoGenerado
	for rows.Next() {
		var reporte entities.ReporteMantenimientoGenerado
		var fechaDesde, fechaHasta, createdAt time.Time
		err := rows.Scan(
			&reporte.ReporteID,
			&reporte.CoordinadorID,
			&fechaDesde,
			&fechaHasta,
			&reporte.Observaciones,
			&createdAt,
		)
		if err != nil {
			log.Println("Error al escanear el reporte de mantenimiento generado:", err)
			return nil, err
		}
		reporte.SetFechaDesde(fechaDesde)
		reporte.SetFechaHasta(fechaHasta)
		reporte.SetCreatedAt(createdAt)
		reportes = append(reportes, reporte)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return reportes, nil
}