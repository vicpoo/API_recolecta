// postgres_anomalia_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type PostgresAnomaliaRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresAnomaliaRepository() repositories.IAnomalia {
	pool := core.GetBD()
	return &PostgresAnomaliaRepository{pool: pool}
}

func (pg *PostgresAnomaliaRepository) Save(anomalia *entities.Anomalia) error {
	query := `
		INSERT INTO anomalia (
			punto_id, 
			tipo_anomalia, 
			descripcion, 
			fecha_reporte, 
			estado,
			fecha_resolucion,
			id_chofer_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING anomalia_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		anomalia.PuntoID, 
		anomalia.TipoAnomalia,
		anomalia.Descripcion,
		anomalia.FechaReporte,
		anomalia.Estado,
		anomalia.FechaResolucion,
		anomalia.IDChoferID,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar la anomalía:", err)
		return err
	}
	
	anomalia.AnomaliaID = id
	return nil
}

func (pg *PostgresAnomaliaRepository) Update(anomalia *entities.Anomalia) error {
	query := `
		UPDATE anomalia
		SET 
			punto_id = $1, 
			tipo_anomalia = $2, 
			descripcion = $3, 
			fecha_reporte = $4, 
			estado = $5,
			fecha_resolucion = $6,
			id_chofer_id = $7
		WHERE anomalia_id = $8
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		anomalia.PuntoID, 
		anomalia.TipoAnomalia,
		anomalia.Descripcion,
		anomalia.FechaReporte,
		anomalia.Estado,
		anomalia.FechaResolucion,
		anomalia.IDChoferID,
		anomalia.AnomaliaID,
	)
	
	if err != nil {
		log.Println("Error al actualizar la anomalía:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("anomalía con ID %d no encontrada", anomalia.AnomaliaID)
	}

	return nil
}

func (pg *PostgresAnomaliaRepository) Delete(id int32) error {
	// Borrado físico
	query := `
		DELETE FROM anomalia
		WHERE anomalia_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar la anomalía:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("anomalía con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresAnomaliaRepository) GetByID(id int32) (*entities.Anomalia, error) {
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		WHERE anomalia_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var anomalia entities.Anomalia
	err := row.Scan(
		&anomalia.AnomaliaID,
		&anomalia.PuntoID,
		&anomalia.TipoAnomalia,
		&anomalia.Descripcion,
		&anomalia.FechaReporte,
		&anomalia.Estado,
		&anomalia.FechaResolucion,
		&anomalia.IDChoferID,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("anomalía con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la anomalía por ID:", err)
		return nil, err
	}

	return &anomalia, nil
}

func (pg *PostgresAnomaliaRepository) GetAll() ([]entities.Anomalia, error) {
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todas las anomalías:", err)
		return nil, err
	}
	defer rows.Close()

	var anomalias []entities.Anomalia
	for rows.Next() {
		var anomalia entities.Anomalia
		err := rows.Scan(
			&anomalia.AnomaliaID,
			&anomalia.PuntoID,
			&anomalia.TipoAnomalia,
			&anomalia.Descripcion,
			&anomalia.FechaReporte,
			&anomalia.Estado,
			&anomalia.FechaResolucion,
			&anomalia.IDChoferID,
		)
		if err != nil {
			log.Println("Error al escanear la anomalía:", err)
			return nil, err
		}
		anomalias = append(anomalias, anomalia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return anomalias, nil
}

func (pg *PostgresAnomaliaRepository) GetByPuntoID(puntoID int32) ([]entities.Anomalia, error) {
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		WHERE punto_id = $1
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, puntoID)
	if err != nil {
		log.Println("Error al obtener anomalías por punto ID:", err)
		return nil, err
	}
	defer rows.Close()

	var anomalias []entities.Anomalia
	for rows.Next() {
		var anomalia entities.Anomalia
		err := rows.Scan(
			&anomalia.AnomaliaID,
			&anomalia.PuntoID,
			&anomalia.TipoAnomalia,
			&anomalia.Descripcion,
			&anomalia.FechaReporte,
			&anomalia.Estado,
			&anomalia.FechaResolucion,
			&anomalia.IDChoferID,
		)
		if err != nil {
			log.Println("Error al escanear la anomalía:", err)
			return nil, err
		}
		anomalias = append(anomalias, anomalia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return anomalias, nil
}

func (pg *PostgresAnomaliaRepository) GetByChoferID(choferID int32) ([]entities.Anomalia, error) {
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		WHERE id_chofer_id = $1
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, choferID)
	if err != nil {
		log.Println("Error al obtener anomalías por chofer ID:", err)
		return nil, err
	}
	defer rows.Close()

	var anomalias []entities.Anomalia
	for rows.Next() {
		var anomalia entities.Anomalia
		err := rows.Scan(
			&anomalia.AnomaliaID,
			&anomalia.PuntoID,
			&anomalia.TipoAnomalia,
			&anomalia.Descripcion,
			&anomalia.FechaReporte,
			&anomalia.Estado,
			&anomalia.FechaResolucion,
			&anomalia.IDChoferID,
		)
		if err != nil {
			log.Println("Error al escanear la anomalía:", err)
			return nil, err
		}
		anomalias = append(anomalias, anomalia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return anomalias, nil
}

func (pg *PostgresAnomaliaRepository) GetByEstado(estado string) ([]entities.Anomalia, error) {
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		WHERE estado = $1
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, estado)
	if err != nil {
		log.Println("Error al obtener anomalías por estado:", err)
		return nil, err
	}
	defer rows.Close()

	var anomalias []entities.Anomalia
	for rows.Next() {
		var anomalia entities.Anomalia
		err := rows.Scan(
			&anomalia.AnomaliaID,
			&anomalia.PuntoID,
			&anomalia.TipoAnomalia,
			&anomalia.Descripcion,
			&anomalia.FechaReporte,
			&anomalia.Estado,
			&anomalia.FechaResolucion,
			&anomalia.IDChoferID,
		)
		if err != nil {
			log.Println("Error al escanear la anomalía:", err)
			return nil, err
		}
		anomalias = append(anomalias, anomalia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return anomalias, nil
}

func (pg *PostgresAnomaliaRepository) GetByTipoAnomalia(tipoAnomalia string) ([]entities.Anomalia, error) {
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		WHERE tipo_anomalia = $1
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, tipoAnomalia)
	if err != nil {
		log.Println("Error al obtener anomalías por tipo:", err)
		return nil, err
	}
	defer rows.Close()

	var anomalias []entities.Anomalia
	for rows.Next() {
		var anomalia entities.Anomalia
		err := rows.Scan(
			&anomalia.AnomaliaID,
			&anomalia.PuntoID,
			&anomalia.TipoAnomalia,
			&anomalia.Descripcion,
			&anomalia.FechaReporte,
			&anomalia.Estado,
			&anomalia.FechaResolucion,
			&anomalia.IDChoferID,
		)
		if err != nil {
			log.Println("Error al escanear la anomalía:", err)
			return nil, err
		}
		anomalias = append(anomalias, anomalia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return anomalias, nil
}

func (pg *PostgresAnomaliaRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Anomalia, error) {
	// Parsear fechas usando la función auxiliar parseFecha
	startTime, err := parseFecha(fechaInicio)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha_inicio inválido: %v", err)
	}
	
	endTime, err := parseFecha(fechaFin)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha_fin inválido: %v", err)
	}
	
	// Si solo se proporcionó fecha (sin hora), ajustar para incluir todo el día final
	if len(fechaFin) <= 10 { // Solo fecha (YYYY-MM-DD)
		endTime = endTime.Add(24 * time.Hour)
	}
	
	query := `
		SELECT 
			anomalia_id,
			punto_id,
			tipo_anomalia,
			descripcion,
			fecha_reporte,
			estado,
			fecha_resolucion,
			id_chofer_id
		FROM anomalia
		WHERE fecha_reporte BETWEEN $1 AND $2
		ORDER BY fecha_reporte DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, startTime, endTime)
	if err != nil {
		log.Println("Error al obtener anomalías por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var anomalias []entities.Anomalia
	for rows.Next() {
		var anomalia entities.Anomalia
		err := rows.Scan(
			&anomalia.AnomaliaID,
			&anomalia.PuntoID,
			&anomalia.TipoAnomalia,
			&anomalia.Descripcion,
			&anomalia.FechaReporte,
			&anomalia.Estado,
			&anomalia.FechaResolucion,
			&anomalia.IDChoferID,
		)
		if err != nil {
			log.Println("Error al escanear la anomalía:", err)
			return nil, err
		}
		anomalias = append(anomalias, anomalia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return anomalias, nil
}