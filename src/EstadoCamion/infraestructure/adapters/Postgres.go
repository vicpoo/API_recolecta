package adapters

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/EstadoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres() *Postgres {
	conn, _ := core.ConnectPostgres()
	return &Postgres{
		conn: conn,
	}
}

//
// CREATE
//
func (pg *Postgres) Save(estado *entities.EstadoCamion) (*entities.EstadoCamion, error) {
	sql := `
	INSERT INTO estado_camion (camion_id, estado, observaciones)
	VALUES ($1, $2, $3)
	RETURNING estado_id, timestamp
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		estado.CamionID,
		estado.Estado,
		estado.Observaciones,
	).Scan(&estado.EstadoID, &estado.Timestamp)

	if err != nil {
		return nil, err
	}

	return estado, nil
}


//
// GET BY ID
//
func (pg *Postgres) GetById(id int32) (*entities.EstadoCamion, error) {
	var estado entities.EstadoCamion

	sql := `
	SELECT 
		estado_id,
		camion_id,
		estado,
		timestamp,
		observaciones
	FROM estado_camion
	WHERE estado_id = $1
	`

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(
		&estado.EstadoID,
		&estado.CamionID,
		&estado.Estado,
		&estado.Timestamp,
		&estado.Observaciones,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("estado_camion no encontrado")
		}
		return nil, err
	}

	return &estado, nil
}

//
// LIST ALL
//
func (pg *Postgres) ListAll() ([]entities.EstadoCamion, error) {
	sql := `
	SELECT 
		estado_id,
		camion_id,
		estado,
		timestamp,
		observaciones
	FROM estado_camion
	ORDER BY estado_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var estados []entities.EstadoCamion

	for rows.Next() {
		var estado entities.EstadoCamion
		err := rows.Scan(
			&estado.EstadoID,
			&estado.CamionID,
			&estado.Estado,
			&estado.Timestamp,
			&estado.Observaciones,
		)
		if err != nil {
			return nil, err
		}

		estados = append(estados, estado)
	}

	return estados, nil
}

//
// UPDATE
//
func (pg *Postgres) Update(id int32, estado *entities.EstadoCamion) (*entities.EstadoCamion, error) {
	sql := `
	UPDATE estado_camion
	SET
		camion_id = $1,
		estado = $2,
		observaciones = $3
	WHERE estado_id = $4
	RETURNING timestamp
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		estado.CamionID,
		estado.Estado,
		estado.Observaciones,
		id,
	).Scan(&estado.Timestamp)

	if err != nil {
		return nil, err
	}

	estado.EstadoID = id
	return estado, nil
}

func (pg *Postgres) Delete(id int32) error {
	sql := `
	DELETE FROM estado_camion
	WHERE estado_id = $1
	`

	_, err := pg.conn.Exec(context.Background(), sql, id)
	return err
}
