// postgres_registro_mantenimiento_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type PostgresRegistroMantenimientoRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRegistroMantenimientoRepository() repositories.IRegistroMantenimiento {
	pool := core.GetBD()
	return &PostgresRegistroMantenimientoRepository{pool: pool}
}

func (pg *PostgresRegistroMantenimientoRepository) Save(registroMantenimiento *entities.RegistroMantenimiento) error {
	query := `
		INSERT INTO registro_mantenimiento (
			alerta_id, 
			camion_id, 
			coordinador_id, 
			mecanico_responsable, 
			fecha_realizada, 
			kilometraje_mantenimiento, 
			observaciones, 
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING registro_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		registroMantenimiento.AlertaID, 
		registroMantenimiento.CamionID,
		registroMantenimiento.CoordinadorID,
		registroMantenimiento.MecanicoResponsable,
		registroMantenimiento.FechaRealizada,
		registroMantenimiento.KilometrajeMantenimiento,
		registroMantenimiento.Observaciones,
		registroMantenimiento.CreatedAt,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar el registro de mantenimiento:", err)
		return err
	}
	
	registroMantenimiento.RegistroID = id
	return nil
}

func (pg *PostgresRegistroMantenimientoRepository) Update(registroMantenimiento *entities.RegistroMantenimiento) error {
	query := `
		UPDATE registro_mantenimiento
		SET 
			alerta_id = $1, 
			camion_id = $2, 
			coordinador_id = $3, 
			mecanico_responsable = $4, 
			fecha_realizada = $5, 
			kilometraje_mantenimiento = $6, 
			observaciones = $7
		WHERE registro_id = $8
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		registroMantenimiento.AlertaID, 
		registroMantenimiento.CamionID,
		registroMantenimiento.CoordinadorID,
		registroMantenimiento.MecanicoResponsable,
		registroMantenimiento.FechaRealizada,
		registroMantenimiento.KilometrajeMantenimiento,
		registroMantenimiento.Observaciones,
		registroMantenimiento.RegistroID,
	)
	
	if err != nil {
		log.Println("Error al actualizar el registro de mantenimiento:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("registro de mantenimiento con ID %d no encontrado", registroMantenimiento.RegistroID)
	}

	return nil
}

func (pg *PostgresRegistroMantenimientoRepository) Delete(id int32) error {
	query := `
		DELETE FROM registro_mantenimiento
		WHERE registro_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar el registro de mantenimiento:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("registro de mantenimiento con ID %d no encontrado", id)
	}

	return nil
}

func (pg *PostgresRegistroMantenimientoRepository) GetByID(id int32) (*entities.RegistroMantenimiento, error) {
	query := `
		SELECT 
			registro_id,
			alerta_id,
			camion_id,
			coordinador_id,
			mecanico_responsable,
			fecha_realizada,
			kilometraje_mantenimiento,
			observaciones,
			created_at
		FROM registro_mantenimiento
		WHERE registro_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var registro entities.RegistroMantenimiento
	err := row.Scan(
		&registro.RegistroID,
		&registro.AlertaID,
		&registro.CamionID,
		&registro.CoordinadorID,
		&registro.MecanicoResponsable,
		&registro.FechaRealizada,
		&registro.KilometrajeMantenimiento,
		&registro.Observaciones,
		&registro.CreatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("registro de mantenimiento con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el registro de mantenimiento por ID:", err)
		return nil, err
	}

	return &registro, nil
}

func (pg *PostgresRegistroMantenimientoRepository) GetAll() ([]entities.RegistroMantenimiento, error) {
	query := `
		SELECT 
			registro_id,
			alerta_id,
			camion_id,
			coordinador_id,
			mecanico_responsable,
			fecha_realizada,
			kilometraje_mantenimiento,
			observaciones,
			created_at
		FROM registro_mantenimiento
		ORDER BY fecha_realizada DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todos los registros de mantenimiento:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroMantenimiento
	for rows.Next() {
		var registro entities.RegistroMantenimiento
		err := rows.Scan(
			&registro.RegistroID,
			&registro.AlertaID,
			&registro.CamionID,
			&registro.CoordinadorID,
			&registro.MecanicoResponsable,
			&registro.FechaRealizada,
			&registro.KilometrajeMantenimiento,
			&registro.Observaciones,
			&registro.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el registro de mantenimiento:", err)
			return nil, err
		}
		registros = append(registros, registro)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return registros, nil
}

func (pg *PostgresRegistroMantenimientoRepository) GetByCamionID(camionID int32) ([]entities.RegistroMantenimiento, error) {
	query := `
		SELECT 
			registro_id,
			alerta_id,
			camion_id,
			coordinador_id,
			mecanico_responsable,
			fecha_realizada,
			kilometraje_mantenimiento,
			observaciones,
			created_at
		FROM registro_mantenimiento
		WHERE camion_id = $1
		ORDER BY fecha_realizada DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, camionID)
	if err != nil {
		log.Println("Error al obtener registros de mantenimiento por camion ID:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroMantenimiento
	for rows.Next() {
		var registro entities.RegistroMantenimiento
		err := rows.Scan(
			&registro.RegistroID,
			&registro.AlertaID,
			&registro.CamionID,
			&registro.CoordinadorID,
			&registro.MecanicoResponsable,
			&registro.FechaRealizada,
			&registro.KilometrajeMantenimiento,
			&registro.Observaciones,
			&registro.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el registro de mantenimiento:", err)
			return nil, err
		}
		registros = append(registros, registro)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return registros, nil
}

func (pg *PostgresRegistroMantenimientoRepository) GetByAlertaID(alertaID int32) (*entities.RegistroMantenimiento, error) {
	query := `
		SELECT 
			registro_id,
			alerta_id,
			camion_id,
			coordinador_id,
			mecanico_responsable,
			fecha_realizada,
			kilometraje_mantenimiento,
			observaciones,
			created_at
		FROM registro_mantenimiento
		WHERE alerta_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, alertaID)

	var registro entities.RegistroMantenimiento
	err := row.Scan(
		&registro.RegistroID,
		&registro.AlertaID,
		&registro.CamionID,
		&registro.CoordinadorID,
		&registro.MecanicoResponsable,
		&registro.FechaRealizada,
		&registro.KilometrajeMantenimiento,
		&registro.Observaciones,
		&registro.CreatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("registro de mantenimiento con alerta ID %d no encontrado", alertaID)
		}
		log.Println("Error al buscar el registro de mantenimiento por alerta ID:", err)
		return nil, err
	}

	return &registro, nil
}

func (pg *PostgresRegistroMantenimientoRepository) GetByCoordinadorID(coordinadorID int32) ([]entities.RegistroMantenimiento, error) {
	query := `
		SELECT 
			registro_id,
			alerta_id,
			camion_id,
			coordinador_id,
			mecanico_responsable,
			fecha_realizada,
			kilometraje_mantenimiento,
			observaciones,
			created_at
		FROM registro_mantenimiento
		WHERE coordinador_id = $1
		ORDER BY fecha_realizada DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, coordinadorID)
	if err != nil {
		log.Println("Error al obtener registros de mantenimiento por coordinador ID:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroMantenimiento
	for rows.Next() {
		var registro entities.RegistroMantenimiento
		err := rows.Scan(
			&registro.RegistroID,
			&registro.AlertaID,
			&registro.CamionID,
			&registro.CoordinadorID,
			&registro.MecanicoResponsable,
			&registro.FechaRealizada,
			&registro.KilometrajeMantenimiento,
			&registro.Observaciones,
			&registro.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el registro de mantenimiento:", err)
			return nil, err
		}
		registros = append(registros, registro)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return registros, nil
}

func (pg *PostgresRegistroMantenimientoRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.RegistroMantenimiento, error) {
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
			registro_id,
			alerta_id,
			camion_id,
			coordinador_id,
			mecanico_responsable,
			fecha_realizada,
			kilometraje_mantenimiento,
			observaciones,
			created_at
		FROM registro_mantenimiento
		WHERE fecha_realizada BETWEEN $1 AND $2
		ORDER BY fecha_realizada DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, startTime, endTime)
	if err != nil {
		log.Println("Error al obtener registros de mantenimiento por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroMantenimiento
	for rows.Next() {
		var registro entities.RegistroMantenimiento
		err := rows.Scan(
			&registro.RegistroID,
			&registro.AlertaID,
			&registro.CamionID,
			&registro.CoordinadorID,
			&registro.MecanicoResponsable,
			&registro.FechaRealizada,
			&registro.KilometrajeMantenimiento,
			&registro.Observaciones,
			&registro.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear el registro de mantenimiento:", err)
			return nil, err
		}
		registros = append(registros, registro)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return registros, nil
}