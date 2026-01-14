package adapters

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres() *Postgres {
	conn, _ := core.ConnectPostgres()
	return &Postgres{conn: conn}
}

//
// CREATE
//
func (pg *Postgres) Save(h *entities.HistorialAsignacionCamion) (*entities.HistorialAsignacionCamion, error) {
	sql := `
	INSERT INTO historial_asignacion_camion (id_chofer, id_camion, fecha_baja, eliminado)
	VALUES ($1, $2, $3, $4)
	RETURNING id_historial, fecha_asignacion, created_at
	`

	err := pg.conn.QueryRow(context.Background(), sql,
		h.IDChofer,
		h.IDCamion,
		h.FechaBaja,
		h.Eliminado,
	).Scan(&h.IDHistorial, &h.FechaAsignacion, &h.CreatedAt)

	if err != nil {
		return nil, err
	}

	return h, nil
}

//
// GET BY ID
//
func (pg *Postgres) GetById(id int32) (*entities.HistorialAsignacionCamion, error) {
	var h entities.HistorialAsignacionCamion

	sql := `
	SELECT id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
	FROM historial_asignacion_camion
	WHERE id_historial = $1
	`

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(
		&h.IDHistorial,
		&h.IDChofer,
		&h.IDCamion,
		&h.FechaAsignacion,
		&h.FechaBaja,
		&h.Eliminado,
		&h.CreatedAt,
		&h.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, errors.New("historial no encontrado")
	}
	if err != nil {
		return nil, err
	}

	return &h, nil
}

//
// LIST ALL
//
func (pg *Postgres) ListAll() ([]entities.HistorialAsignacionCamion, error) {
	sql := `
	SELECT id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
	FROM historial_asignacion_camion
	ORDER BY id_historial DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entities.HistorialAsignacionCamion

	for rows.Next() {
		var h entities.HistorialAsignacionCamion
		err := rows.Scan(
			&h.IDHistorial,
			&h.IDChofer,
			&h.IDCamion,
			&h.FechaAsignacion,
			&h.FechaBaja,
			&h.Eliminado,
			&h.CreatedAt,
			&h.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, h)
	}

	return list, nil
}

//
// UPDATE
//
func (pg *Postgres) Update(id int32, h *entities.HistorialAsignacionCamion) (*entities.HistorialAsignacionCamion, error) {
	sql := `
	UPDATE historial_asignacion_camion
	SET id_chofer=$1, id_camion=$2, fecha_baja=$3, eliminado=$4, updated_at=now()
	WHERE id_historial=$5
	RETURNING fecha_asignacion, created_at, updated_at
	`

	err := pg.conn.QueryRow(context.Background(), sql,
		h.IDChofer,
		h.IDCamion,
		h.FechaBaja,
		h.Eliminado,
		id,
	).Scan(&h.FechaAsignacion, &h.CreatedAt, &h.UpdatedAt)

	if err != nil {
		return nil, err
	}

	h.IDHistorial = int(id)
	return h, nil
}

//
// DELETE (Soft)
//
func (pg *Postgres) Delete(id int32) error {
	_, err := pg.conn.Exec(context.Background(),
		`UPDATE historial_asignacion_camion SET eliminado=true, updated_at=now() WHERE id_historial=$1`, id)
	return err
}

//
// ============ MÉTODOS AVANZADOS ============
//

func (pg *Postgres) GetByCamionId(camionId int32) ([]entities.HistorialAsignacionCamion, error) {
	sql := `
	SELECT id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
	FROM historial_asignacion_camion
	WHERE id_camion=$1
	ORDER BY fecha_asignacion DESC
	`
	return pg.fetchMany(sql, camionId)
}

func (pg *Postgres) GetByChoferId(choferId int32) ([]entities.HistorialAsignacionCamion, error) {
	sql := `
	SELECT id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
	FROM historial_asignacion_camion
	WHERE id_chofer=$1
	ORDER BY fecha_asignacion DESC
	`
	return pg.fetchMany(sql, choferId)
}

func (pg *Postgres) GetActivoByCamionId(camionId int32) (*entities.HistorialAsignacionCamion, error) {
	sql := `
	SELECT id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
	FROM historial_asignacion_camion
	WHERE id_camion=$1 AND fecha_baja IS NULL AND eliminado=false
	LIMIT 1
	`
	return pg.fetchOne(sql, camionId)
}

func (pg *Postgres) GetActivoByChoferId(choferId int32) (*entities.HistorialAsignacionCamion, error) {
	sql := `
	SELECT id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
	FROM historial_asignacion_camion
	WHERE id_chofer=$1 AND fecha_baja IS NULL AND eliminado=false
	LIMIT 1
	`
	return pg.fetchOne(sql, choferId)
}

func (pg *Postgres) DarDeBaja(id int32) error {
	_, err := pg.conn.Exec(context.Background(),
		`UPDATE historial_asignacion_camion SET fecha_baja=now(), updated_at=now() WHERE id_historial=$1`, id)
	return err
}

func (pg *Postgres) CerrarAsignacionActivaCamion(camionId int32) error {
	_, err := pg.conn.Exec(context.Background(),
		`UPDATE historial_asignacion_camion SET fecha_baja=now(), updated_at=now() WHERE id_camion=$1 AND fecha_baja IS NULL`, camionId)
	return err
}

func (pg *Postgres) CerrarAsignacionActivaChofer(choferId int32) error {
	_, err := pg.conn.Exec(context.Background(),
		`UPDATE historial_asignacion_camion SET fecha_baja=now(), updated_at=now() WHERE id_chofer=$1 AND fecha_baja IS NULL`, choferId)
	return err
}

//
// Helpers
//
func (pg *Postgres) fetchOne(sql string, param any) (*entities.HistorialAsignacionCamion, error) {
	var h entities.HistorialAsignacionCamion
	err := pg.conn.QueryRow(context.Background(), sql, param).Scan(
		&h.IDHistorial, &h.IDChofer, &h.IDCamion, &h.FechaAsignacion,
		&h.FechaBaja, &h.Eliminado, &h.CreatedAt, &h.UpdatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("historial_asignacion_camion no encontrado")
		}

		return nil, err
	}
	return &h, nil
}

func (pg *Postgres) fetchMany(sql string, param any) ([]entities.HistorialAsignacionCamion, error) {
	rows, err := pg.conn.Query(context.Background(), sql, param)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entities.HistorialAsignacionCamion
	for rows.Next() {
		var h entities.HistorialAsignacionCamion
		err := rows.Scan(
			&h.IDHistorial, &h.IDChofer, &h.IDCamion, &h.FechaAsignacion,
			&h.FechaBaja, &h.Eliminado, &h.CreatedAt, &h.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, h)
	}
	return list, nil
}
