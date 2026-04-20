// postgres_notificacion_repository.go
package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/core"
	repositories "github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type PostgresNotificacionRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresNotificacionRepository() repositories.INotificacion {
	pool := core.GetBD()
	return &PostgresNotificacionRepository{pool: pool}
}

// ==================== CRUD BÁSICO ====================

func (pg *PostgresNotificacionRepository) Save(notificacion *entities.Notificacion) error {
	query := `
		INSERT INTO notificacion (
			usuario_id, 
			tipo, 
			titulo, 
			mensaje, 
			activa, 
			id_camion_relacionado, 
			id_falla_relacionado, 
			id_mantenimiento_relacionado, 
			creado_por, 
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING notificacion_id
	`
	
	ctx := context.Background()
	
	var id int32
	err := pg.pool.QueryRow(
		ctx,
		query, 
		notificacion.UsuarioID, 
		notificacion.Tipo,
		notificacion.Titulo,
		notificacion.Mensaje,
		notificacion.Activa,
		notificacion.IDCamionRelacionado,
		notificacion.IDFallaRelacionado,
		notificacion.IDMantenimientoRelacionado,
		notificacion.CreadoPor,
		notificacion.CreatedAt,
	).Scan(&id)
	
	if err != nil {
		log.Println("Error al guardar la notificación:", err)
		return err
	}
	
	notificacion.NotificacionID = id
	return nil
}

func (pg *PostgresNotificacionRepository) SaveForMultipleUsers(notificacion *entities.Notificacion, usuarioIDs []int32) error {
	if len(usuarioIDs) == 0 {
		return fmt.Errorf("se requiere al menos un usuario destinatario")
	}

	ctx := context.Background()
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		log.Println("Error al iniciar transacción:", err)
		return err
	}
	defer tx.Rollback(ctx)

	var notificacionID int32

	for _, usuarioID := range usuarioIDs {
		query := `
			INSERT INTO notificacion (
				usuario_id, tipo, titulo, mensaje, activa, 
				id_camion_relacionado, id_falla_relacionado, 
				id_mantenimiento_relacionado, creado_por, created_at
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			RETURNING notificacion_id
		`

		err = tx.QueryRow(
			ctx, query,
			usuarioID,
			notificacion.Tipo,
			notificacion.Titulo,
			notificacion.Mensaje,
			notificacion.Activa,
			notificacion.IDCamionRelacionado,
			notificacion.IDFallaRelacionado,
			notificacion.IDMantenimientoRelacionado,
			notificacion.CreadoPor,
			notificacion.CreatedAt,
		).Scan(&notificacionID)

		if err != nil {
			log.Printf("Error al guardar notificación para usuario %d: %v", usuarioID, err)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Println("Error al hacer commit:", err)
		return err
	}

	log.Printf("Notificaciones creadas para %d usuarios", len(usuarioIDs))
	return nil
}

func (pg *PostgresNotificacionRepository) SaveForAllUsers(notificacion *entities.Notificacion) error {
	queryGetUsers := `SELECT user_id FROM usuario WHERE eliminado = false OR eliminado IS NULL`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, queryGetUsers)
	if err != nil {
		log.Println("Error al obtener usuarios:", err)
		return err
	}
	defer rows.Close()

	var usuarioIDs []int32
	for rows.Next() {
		var userID int32
		if err := rows.Scan(&userID); err != nil {
			log.Println("Error al escanear usuario:", err)
			return err
		}
		usuarioIDs = append(usuarioIDs, userID)
	}

	if len(usuarioIDs) == 0 {
		return fmt.Errorf("no hay usuarios disponibles para notificar")
	}

	return pg.SaveForMultipleUsers(notificacion, usuarioIDs)
}

func (pg *PostgresNotificacionRepository) Update(notificacion *entities.Notificacion) error {
	query := `
		UPDATE notificacion
		SET 
			usuario_id = $1, 
			tipo = $2, 
			titulo = $3, 
			mensaje = $4, 
			activa = $5, 
			id_camion_relacionado = $6, 
			id_falla_relacionado = $7, 
			id_mantenimiento_relacionado = $8, 
			creado_por = $9
		WHERE notificacion_id = $10
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(
		ctx,
		query, 
		notificacion.UsuarioID, 
		notificacion.Tipo,
		notificacion.Titulo,
		notificacion.Mensaje,
		notificacion.Activa,
		notificacion.IDCamionRelacionado,
		notificacion.IDFallaRelacionado,
		notificacion.IDMantenimientoRelacionado,
		notificacion.CreadoPor,
		notificacion.NotificacionID,
	)
	
	if err != nil {
		log.Println("Error al actualizar la notificación:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("notificación con ID %d no encontrada", notificacion.NotificacionID)
	}

	return nil
}

func (pg *PostgresNotificacionRepository) Delete(id int32) error {
	query := `
		DELETE FROM notificacion
		WHERE notificacion_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al eliminar la notificación:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("notificación con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresNotificacionRepository) GetByID(id int32) (*entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE notificacion_id = $1
	`
	
	ctx := context.Background()
	row := pg.pool.QueryRow(ctx, query, id)

	var notificacion entities.Notificacion
	err := row.Scan(
		&notificacion.NotificacionID,
		&notificacion.UsuarioID,
		&notificacion.Tipo,
		&notificacion.Titulo,
		&notificacion.Mensaje,
		&notificacion.Activa,
		&notificacion.IDCamionRelacionado,
		&notificacion.IDFallaRelacionado,
		&notificacion.IDMantenimientoRelacionado,
		&notificacion.CreadoPor,
		&notificacion.CreatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("notificación con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la notificación por ID:", err)
		return nil, err
	}

	return &notificacion, nil
}

func (pg *PostgresNotificacionRepository) GetAll() ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener todas las notificaciones:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetActivas() ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE activa = true
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener notificaciones activas:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetInactivas() ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE activa = false
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener notificaciones inactivas:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetGlobales() ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE usuario_id IS NULL
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		log.Println("Error al obtener notificaciones globales:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByUsuarioID(usuarioID int32) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE usuario_id = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, usuarioID)
	if err != nil {
		log.Println("Error al obtener notificaciones por usuario ID:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetActivasByUsuarioID(usuarioID int32) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE usuario_id = $1 AND activa = true
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, usuarioID)
	if err != nil {
		log.Println("Error al obtener notificaciones activas por usuario ID:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByTipo(tipo string) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE tipo = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, tipo)
	if err != nil {
		log.Println("Error al obtener notificaciones por tipo:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByUsuarioYTipo(usuarioID int32, tipo string) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE usuario_id = $1 AND tipo = $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, usuarioID, tipo)
	if err != nil {
		log.Println("Error al obtener notificaciones por usuario y tipo:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByCamionID(camionID int32) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE id_camion_relacionado = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, camionID)
	if err != nil {
		log.Println("Error al obtener notificaciones por camión ID:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByCamionYTipo(camionID int32, tipo string) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE id_camion_relacionado = $1 AND tipo = $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, camionID, tipo)
	if err != nil {
		log.Println("Error al obtener notificaciones por camión y tipo:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByCreadoPor(creadoPor int32) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE creado_por = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, creadoPor)
	if err != nil {
		log.Println("Error al obtener notificaciones por creador:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByFallaID(fallaID int32) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE id_falla_relacionado = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, fallaID)
	if err != nil {
		log.Println("Error al obtener notificaciones por falla ID:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByMantenimientoID(mantenimientoID int32) ([]entities.Notificacion, error) {
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE id_mantenimiento_relacionado = $1
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, mantenimientoID)
	if err != nil {
		log.Println("Error al obtener notificaciones por mantenimiento ID:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) GetByFechaRange(fechaInicio, fechaFin string) ([]entities.Notificacion, error) {
	startTime, err := time.Parse("2006-01-02", fechaInicio)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha_inicio inválido: %v", err)
	}
	
	endTime, err := time.Parse("2006-01-02", fechaFin)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha_fin inválido: %v", err)
	}
	
	endTime = endTime.Add(24 * time.Hour)
	
	query := `
		SELECT 
			notificacion_id,
			usuario_id,
			tipo,
			titulo,
			mensaje,
			activa,
			id_camion_relacionado,
			id_falla_relacionado,
			id_mantenimiento_relacionado,
			creado_por,
			created_at
		FROM notificacion
		WHERE created_at BETWEEN $1 AND $2
		ORDER BY created_at DESC
	`
	
	ctx := context.Background()
	rows, err := pg.pool.Query(ctx, query, startTime, endTime)
	if err != nil {
		log.Println("Error al obtener notificaciones por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var notificaciones []entities.Notificacion
	for rows.Next() {
		var notificacion entities.Notificacion
		err := rows.Scan(
			&notificacion.NotificacionID,
			&notificacion.UsuarioID,
			&notificacion.Tipo,
			&notificacion.Titulo,
			&notificacion.Mensaje,
			&notificacion.Activa,
			&notificacion.IDCamionRelacionado,
			&notificacion.IDFallaRelacionado,
			&notificacion.IDMantenimientoRelacionado,
			&notificacion.CreadoPor,
			&notificacion.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear la notificación:", err)
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return notificaciones, nil
}

func (pg *PostgresNotificacionRepository) CountByUsuarioID(usuarioID int32) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM notificacion
		WHERE usuario_id = $1
	`
	
	ctx := context.Background()
	var count int64
	err := pg.pool.QueryRow(ctx, query, usuarioID).Scan(&count)
	if err != nil {
		log.Println("Error al contar notificaciones por usuario ID:", err)
		return 0, err
	}

	return count, nil
}

func (pg *PostgresNotificacionRepository) CountActivasByUsuarioID(usuarioID int32) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM notificacion
		WHERE usuario_id = $1 AND activa = true
	`
	
	ctx := context.Background()
	var count int64
	err := pg.pool.QueryRow(ctx, query, usuarioID).Scan(&count)
	if err != nil {
		log.Println("Error al contar notificaciones activas por usuario ID:", err)
		return 0, err
	}

	return count, nil
}

func (pg *PostgresNotificacionRepository) CountByTipo(tipo string) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM notificacion
		WHERE tipo = $1
	`
	
	ctx := context.Background()
	var count int64
	err := pg.pool.QueryRow(ctx, query, tipo).Scan(&count)
	if err != nil {
		log.Println("Error al contar notificaciones por tipo:", err)
		return 0, err
	}

	return count, nil
}

func (pg *PostgresNotificacionRepository) CountByCamionID(camionID int32) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM notificacion
		WHERE id_camion_relacionado = $1
	`
	
	ctx := context.Background()
	var count int64
	err := pg.pool.QueryRow(ctx, query, camionID).Scan(&count)
	if err != nil {
		log.Println("Error al contar notificaciones por camión ID:", err)
		return 0, err
	}

	return count, nil
}

func (pg *PostgresNotificacionRepository) MarcarComoLeida(id int32) error {
	query := `
		UPDATE notificacion
		SET activa = false
		WHERE notificacion_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al marcar notificación como leída:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("notificación con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresNotificacionRepository) MarcarComoActiva(id int32) error {
	query := `
		UPDATE notificacion
		SET activa = true
		WHERE notificacion_id = $1
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error al marcar notificación como activa:", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("notificación con ID %d no encontrada", id)
	}

	return nil
}

func (pg *PostgresNotificacionRepository) MarcarTodasComoLeidas(usuarioID int32) error {
	query := `
		UPDATE notificacion
		SET activa = false
		WHERE usuario_id = $1 AND activa = true
	`
	
	ctx := context.Background()
	cmdTag, err := pg.pool.Exec(ctx, query, usuarioID)
	if err != nil {
		log.Println("Error al marcar todas las notificaciones como leídas:", err)
		return err
	}

	log.Printf("Notificaciones marcadas como leídas: %d filas afectadas", cmdTag.RowsAffected())
	return nil
}

func (pg *PostgresNotificacionRepository) CrearNotificacionFalla(usuarioID *int32, titulo string, mensaje string, camionID int32, fallaID int32, creadoPor *int32) error {
	notificacion := entities.NewNotificacionFalla(
		usuarioID,
		titulo,
		mensaje,
		camionID,
		fallaID,
		creadoPor,
	)
	
	return pg.Save(notificacion)
}

func (pg *PostgresNotificacionRepository) CrearNotificacionMantenimiento(usuarioID *int32, titulo string, mensaje string, camionID int32, mantenimientoID int32, creadoPor *int32) error {
	notificacion := entities.NewNotificacionMantenimiento(
		usuarioID,
		titulo,
		mensaje,
		camionID,
		mantenimientoID,
		creadoPor,
	)
	
	return pg.Save(notificacion)
}

func (pg *PostgresNotificacionRepository) CrearNotificacionEmergencia(usuarioID *int32, titulo string, mensaje string, camionID int32, creadoPor *int32) error {
	notificacion := entities.NewNotificacionEmergencia(
		usuarioID,
		titulo,
		mensaje,
		camionID,
		creadoPor,
	)
	
	return pg.Save(notificacion)
}