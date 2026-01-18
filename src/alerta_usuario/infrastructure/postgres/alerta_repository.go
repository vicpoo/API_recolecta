package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"
)

type AlertaRepository struct {
	db *pgxpool.Pool
}

func NewAlertaRepository(db *pgxpool.Pool) *AlertaRepository {
	return &AlertaRepository{db}
}

func (r *AlertaRepository) Create(a *domain.AlertaUsuario) error {
	query := `
		INSERT INTO alerta_usuario
		(titulo, mensaje, usuario_id, creado_por, leida, created_at)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		a.Titulo,
		a.Mensaje,
		a.UsuarioID,
		a.CreadoPor,
		a.Leida,
		a.CreatedAt,
	)

	return err
}

func (r *AlertaRepository) GetByUsuario(usuarioID int) ([]domain.AlertaUsuario, error) {
	query := `
		SELECT alerta_id, titulo, mensaje, leida, created_at, creado_por
		FROM alerta_usuario
		WHERE usuario_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(context.Background(), query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alertas []domain.AlertaUsuario

	for rows.Next() {
		var a domain.AlertaUsuario
		if err := rows.Scan(
			&a.AlertaID,
			&a.Titulo,
			&a.Mensaje,
			&a.Leida,
			&a.CreatedAt,
			&a.CreadoPor,
		); err != nil {
			return nil, err
		}
		alertas = append(alertas, a)
	}

	return alertas, nil
}

func (r *AlertaRepository) MarkAsRead(alertaID int, usuarioID int) error {
	query := `
		UPDATE alerta_usuario
		SET leida = true
		WHERE alerta_id = $1 AND usuario_id = $2
	`

	_, err := r.db.Exec(context.Background(), query, alertaID, usuarioID)
	return err
}
