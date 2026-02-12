package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type PostgresRutaCamion struct {
	conn *pgxpool.Pool
}

func NewPostgresRutaCamion() *PostgresRutaCamion {
	conn, _ := core.ConnectPostgres()
	return &PostgresRutaCamion{
		conn: conn,
	}
}

//
// CREATE
//
func (pg *PostgresRutaCamion) Save(rutaCamion *entities.RutaCamion) (*entities.RutaCamion, error) {
	rutaCamion.CreatedAt = time.Now()
	sql := `
	INSERT INTO ruta_camion
	(
		ruta_id,
		camion_id,
		fecha,
		created_at
	)
	VALUES ($1, $2, $3, $4)
	RETURNING ruta_camion_id
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		rutaCamion.RutaID,
		rutaCamion.CamionID,
		rutaCamion.Fecha,      // dato de negocio
		rutaCamion.CreatedAt,  // 👈 tú lo insertas
	).Scan(&rutaCamion.RutaCamionID)

	if err != nil {
		return nil, err
	}

	return rutaCamion, nil
}


//
// UPDATE
//
func (pg *PostgresRutaCamion) Update(id int32, rutaCamion *entities.RutaCamion) (*entities.RutaCamion, error) {
	sql := `
	UPDATE ruta_camion
	SET
		ruta_id = $1,
		camion_id = $2,
		fecha = $3
	WHERE ruta_camion_id = $4
	  AND eliminado = false
	`

	cmd, err := pg.conn.Exec(
		context.Background(),
		sql,
		rutaCamion.RutaID,
		rutaCamion.CamionID,
		rutaCamion.Fecha,
		id,
	)

	if err != nil {
		return nil, err
	}

	if cmd.RowsAffected() == 0 {
		return nil, errors.New("ruta_camion no encontrada")
	}

	rutaCamion.RutaCamionID = id
	return rutaCamion, nil
}


//
// GET ALL
//
func (pg *PostgresRutaCamion) ListAll() ([]entities.RutaCamion, error) {
	sql := `
	SELECT
		ruta_camion_id,
		ruta_id,
		camion_id,
		fecha,
		created_at,
		eliminado
	FROM ruta_camion
	WHERE eliminado = false
	ORDER BY ruta_camion_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rutas []entities.RutaCamion

	for rows.Next() {
		var r entities.RutaCamion
		var fecha time.Time

		err := rows.Scan(
			&r.RutaCamionID,
			&r.RutaID,
			&r.CamionID,
			&fecha,
			&r.CreatedAt,
			&r.Eliminado,
		)
		if err != nil {
			return nil, err
		}

		r.Fecha = fecha
		rutas = append(rutas, r)
	}

	return rutas, nil
}

//
// GET BY ID
//
func (pg *PostgresRutaCamion) GetByID(id int32) (*entities.RutaCamion, error) {
	var r entities.RutaCamion
	var fecha time.Time

	sql := `
	SELECT
		ruta_camion_id,
		ruta_id,
		camion_id,
		fecha,
		created_at,
		eliminado
	FROM ruta_camion
	WHERE ruta_camion_id = $1
	  AND eliminado = false
	`

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(
		&r.RutaCamionID,
		&r.RutaID,
		&r.CamionID,
		&fecha,
		&r.CreatedAt,
		&r.Eliminado,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("ruta_camion no encontrada")
		}
		return nil, err
	}

	r.Fecha = fecha
	return &r, nil
}

//
// GET BY CAMION ID
//
func (pg *PostgresRutaCamion) GetByCamionID(camionID int32) ([]entities.RutaCamion, error) {
	sql := `
	SELECT
		ruta_camion_id,
		ruta_id,
		camion_id,
		fecha,
		created_at,
		eliminado
	FROM ruta_camion
	WHERE camion_id = $1
	  AND eliminado = false
	ORDER BY fecha DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql, camionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rutas []entities.RutaCamion

	for rows.Next() {
		var r entities.RutaCamion
		var fecha time.Time

		err := rows.Scan(
			&r.RutaCamionID,
			&r.RutaID,
			&r.CamionID,
			&fecha,
			&r.CreatedAt,
			&r.Eliminado,
		)
		if err != nil {
			return nil, err
		}

		r.Fecha = fecha
		rutas = append(rutas, r)
	}

	return rutas, nil
}

//
// GET BY RUTA ID
//
func (pg *PostgresRutaCamion) GetByRutaID(rutaID int32) ([]entities.RutaCamion, error) {
	sql := `
	SELECT
		ruta_camion_id,
		ruta_id,
		camion_id,
		fecha,
		created_at,
		eliminado
	FROM ruta_camion
	WHERE ruta_id = $1
	  AND eliminado = false
	ORDER BY fecha DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql, rutaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rutas []entities.RutaCamion

	for rows.Next() {
		var r entities.RutaCamion
		var fecha time.Time

		err := rows.Scan(
			&r.RutaCamionID,
			&r.RutaID,
			&r.CamionID,
			&fecha,
			&r.CreatedAt,
			&r.Eliminado,
		)
		if err != nil {
			return nil, err
		}

		r.Fecha = fecha
		rutas = append(rutas, r)
	}

	return rutas, nil
}

//
// EXISTS BY ID
//
func (pg *PostgresRutaCamion) ExistsByID(id int32) (bool, error) {
	sql := `
	SELECT EXISTS (
		SELECT 1
		FROM ruta_camion
		WHERE ruta_camion_id = $1
		  AND eliminado = false
	)
	`

	var exists bool
	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(&exists)
	return exists, err
}

//
// DELETE (LÓGICO)
//
func (pg *PostgresRutaCamion) Delete(id int32) error {
	sql := `
	UPDATE ruta_camion
	SET eliminado = true
	WHERE ruta_camion_id = $1
	`

	cmd, err := pg.conn.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("ruta_camion no encontrada")
	}

	return nil
}