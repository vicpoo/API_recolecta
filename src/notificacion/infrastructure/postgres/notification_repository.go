package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type NotificacionRepository struct {
	db *pgxpool.Pool
}

func NewNotificacionRepository(db *pgxpool.Pool) *NotificacionRepository {
	return &NotificacionRepository{db}
}

func (r *NotificacionRepository) Create(n *domain.Notificacion) error {
	query := `
		INSERT INTO notificacion
		(usuario_id, tipo, titulo, mensaje, activa,
		 id_camion_relacionado, id_falla_relacionado,
		 id_mantenimiento_relacionado, creado_por, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		n.UsuarioID,
		n.Tipo,
		n.Titulo,
		n.Mensaje,
		n.Activa,
		n.IDCamionRelacionado,
		n.IDFallaRelacionado,
		n.IDMantenimientoRelacionado,
		n.CreadoPor,
		n.CreatedAt,
	)

	return err
}

func (r *NotificacionRepository) GetByID(id int) (*domain.Notificacion, error) {
	query := `
		SELECT notificacion_id, usuario_id, tipo, titulo, mensaje, activa,
		       id_camion_relacionado, id_falla_relacionado,
		       id_mantenimiento_relacionado, creado_por, created_at
		FROM notificacion
		WHERE notificacion_id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)

	var n domain.Notificacion
	err := row.Scan(
		&n.NotificacionID,
		&n.UsuarioID,
		&n.Tipo,
		&n.Titulo,
		&n.Mensaje,
		&n.Activa,
		&n.IDCamionRelacionado,
		&n.IDFallaRelacionado,
		&n.IDMantenimientoRelacionado,
		&n.CreadoPor,
		&n.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &n, nil
}

func (r *NotificacionRepository) GetByUsuario(usuarioID int) ([]domain.Notificacion, error) {
	query := `
		SELECT notificacion_id, usuario_id, tipo, titulo, mensaje, activa,
		       id_camion_relacionado, id_falla_relacionado,
		       id_mantenimiento_relacionado, creado_por, created_at
		FROM notificacion
		WHERE usuario_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(context.Background(), query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notificaciones []domain.Notificacion

	for rows.Next() {
		var n domain.Notificacion
		if err := rows.Scan(
			&n.NotificacionID,
			&n.UsuarioID,
			&n.Tipo,
			&n.Titulo,
			&n.Mensaje,
			&n.Activa,
			&n.IDCamionRelacionado,
			&n.IDFallaRelacionado,
			&n.IDMantenimientoRelacionado,
			&n.CreadoPor,
			&n.CreatedAt,
		); err != nil {
			return nil, err
		}

		notificaciones = append(notificaciones, n)
	}

	return notificaciones, nil
}

func (r *NotificacionRepository) Deactivate(id int) error {
	query := `
		UPDATE notificacion
		SET activa = false
		WHERE notificacion_id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}
