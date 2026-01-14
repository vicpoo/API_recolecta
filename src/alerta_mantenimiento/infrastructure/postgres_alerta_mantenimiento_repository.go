// postgres_alerta_mantenimiento_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/domain/entities"
)

type PostgresAlertaMantenimientoRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresAlertaMantenimientoRepository() repositories.IAlertaMantenimiento {
	pool := core.GetBD()
	return &PostgresAlertaMantenimientoRepository{pool: pool}
}

func (pg *PostgresAlertaMantenimientoRepository) Save(alertaMantenimiento *entities.AlertaMantenimiento) error {
	query := `
		INSERT INTO alerta_mantenimiento (camion_id, tipo_mantenimiento_id, descripcion, observaciones, created_at, atendido)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING alerta_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		alertaMantenimiento.CamionID,
		alertaMantenimiento.TipoMantenimientoID,
		alertaMantenimiento.Descripcion,
		alertaMantenimiento.Observaciones,
		alertaMantenimiento.CreatedAt,
		alertaMantenimiento.Atendido,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar la alerta de mantenimiento:", err)
		return err
	}
	
	alertaMantenimiento.AlertaID = id
	return nil
}

func (pg *PostgresAlertaMantenimientoRepository) Update(alertaMantenimiento *entities.AlertaMantenimiento) error {
	query := `
		UPDATE alerta_mantenimiento
		SET camion_id = $1, tipo_mantenimiento_id = $2, descripcion = $3, 
			observaciones = $4, atendido = $5
		WHERE alerta_id = $6
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		alertaMantenimiento.CamionID,
		alertaMantenimiento.TipoMantenimientoID,
		alertaMantenimiento.Descripcion,
		alertaMantenimiento.Observaciones,
		alertaMantenimiento.Atendido,
		alertaMantenimiento.AlertaID,
	)
	
	if err != nil {
		log.Println("Error al actualizar la alerta de mantenimiento:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("alerta de mantenimiento con ID %d no encontrada", alertaMantenimiento.AlertaID)
	}

	return nil
}

func (pg *PostgresAlertaMantenimientoRepository) Delete(id int32) error {
	query := "DELETE FROM alerta_mantenimiento WHERE alerta_id = $1"
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar la alerta de mantenimiento:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("alerta de mantenimiento con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetAll() ([]entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todas las alertas de mantenimiento:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []entities.AlertaMantenimiento
	for rows.Next() {
		var alerta entities.AlertaMantenimiento
		err := rows.Scan(
			&alerta.AlertaID,
			&alerta.CamionID,
			&alerta.TipoMantenimientoID,
			&alerta.Descripcion,
			&alerta.Observaciones,
			&alerta.CreatedAt,
			&alerta.Atendido,
		)
		if err != nil {
			log.Println("Error al escanear la alerta de mantenimiento:", err)
			return nil, err
		}
		alertas = append(alertas, alerta)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return alertas, nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetByID(id int32) (*entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		WHERE alerta_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var alerta entities.AlertaMantenimiento
	err := row.Scan(
		&alerta.AlertaID,
		&alerta.CamionID,
		&alerta.TipoMantenimientoID,
		&alerta.Descripcion,
		&alerta.Observaciones,
		&alerta.CreatedAt,
		&alerta.Atendido,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("alerta de mantenimiento con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la alerta de mantenimiento por ID:", err)
		return nil, err
	}

	return &alerta, nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetByCamionID(camionID int32) ([]entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		WHERE camion_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, camionID)
	if err != nil {
		log.Println("Error al obtener alertas por camión:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []entities.AlertaMantenimiento
	for rows.Next() {
		var alerta entities.AlertaMantenimiento
		err := rows.Scan(
			&alerta.AlertaID,
			&alerta.CamionID,
			&alerta.TipoMantenimientoID,
			&alerta.Descripcion,
			&alerta.Observaciones,
			&alerta.CreatedAt,
			&alerta.Atendido,
		)
		if err != nil {
			log.Println("Error al escanear la alerta de mantenimiento:", err)
			return nil, err
		}
		alertas = append(alertas, alerta)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return alertas, nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetByTipoMantenimientoID(tipoID int32) ([]entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		WHERE tipo_mantenimiento_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, tipoID)
	if err != nil {
		log.Println("Error al obtener alertas por tipo de mantenimiento:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []entities.AlertaMantenimiento
	for rows.Next() {
		var alerta entities.AlertaMantenimiento
		err := rows.Scan(
			&alerta.AlertaID,
			&alerta.CamionID,
			&alerta.TipoMantenimientoID,
			&alerta.Descripcion,
			&alerta.Observaciones,
			&alerta.CreatedAt,
			&alerta.Atendido,
		)
		if err != nil {
			log.Println("Error al escanear la alerta de mantenimiento:", err)
			return nil, err
		}
		alertas = append(alertas, alerta)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return alertas, nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetPendientes() ([]entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		WHERE atendido = false
		ORDER BY created_at ASC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener alertas pendientes:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []entities.AlertaMantenimiento
	for rows.Next() {
		var alerta entities.AlertaMantenimiento
		err := rows.Scan(
			&alerta.AlertaID,
			&alerta.CamionID,
			&alerta.TipoMantenimientoID,
			&alerta.Descripcion,
			&alerta.Observaciones,
			&alerta.CreatedAt,
			&alerta.Atendido,
		)
		if err != nil {
			log.Println("Error al escanear la alerta de mantenimiento:", err)
			return nil, err
		}
		alertas = append(alertas, alerta)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return alertas, nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetAtendidas() ([]entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		WHERE atendido = true
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener alertas atendidas:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []entities.AlertaMantenimiento
	for rows.Next() {
		var alerta entities.AlertaMantenimiento
		err := rows.Scan(
			&alerta.AlertaID,
			&alerta.CamionID,
			&alerta.TipoMantenimientoID,
			&alerta.Descripcion,
			&alerta.Observaciones,
			&alerta.CreatedAt,
			&alerta.Atendido,
		)
		if err != nil {
			log.Println("Error al escanear la alerta de mantenimiento:", err)
			return nil, err
		}
		alertas = append(alertas, alerta)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return alertas, nil
}

func (pg *PostgresAlertaMantenimientoRepository) MarcarComoAtendida(id int32) error {
	query := `
		UPDATE alerta_mantenimiento
		SET atendido = true
		WHERE alerta_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al marcar alerta como atendida:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("alerta de mantenimiento con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresAlertaMantenimientoRepository) GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.AlertaMantenimiento, error) {
	query := `
		SELECT alerta_id, camion_id, tipo_mantenimiento_id, descripcion, 
			   observaciones, created_at, atendido
		FROM alerta_mantenimiento
		WHERE created_at BETWEEN $1 AND $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener alertas por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []entities.AlertaMantenimiento
	for rows.Next() {
		var alerta entities.AlertaMantenimiento
		err := rows.Scan(
			&alerta.AlertaID,
			&alerta.CamionID,
			&alerta.TipoMantenimientoID,
			&alerta.Descripcion,
			&alerta.Observaciones,
			&alerta.CreatedAt,
			&alerta.Atendido,
		)
		if err != nil {
			log.Println("Error al escanear la alerta de mantenimiento:", err)
			return nil, err
		}
		alertas = append(alertas, alerta)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return alertas, nil
}