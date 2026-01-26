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

type Postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres() *Postgres {
	conn,_ := core.ConnectPostgres()
	return &Postgres{
		conn: conn,
	}
}

func (pg *Postgres) Save(camion *entities.Camion) (*entities.Camion, error) {
	camion.CreatedAt = time.Now()
	sql := `
	INSERT INTO camion
	(
		placa,
		modelo,
		tipo_camion_id,
		es_rentado,
		nombre_disponibilidad,
		color_disponibilidad,
		created_at,
		updated_at
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7, NULL)
	RETURNING camion_id, disponibilidad_id
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		camion.Placa,
		camion.Modelo,
		camion.TipoCamionID,
		camion.EsRentado,
		camion.NombreDisponibilidad,
		camion.ColorDisponibilidad,
		camion.CreatedAt,
	).Scan(&camion.CamionID, &camion.DisponibilidadID)

	if err != nil {
		return nil, err
	}

	return camion, nil
}


func (pg *Postgres) ListAll() ([]entities.Camion, error) {
	rows, err := pg.conn.Query(context.Background(),
		`SELECT camion_id, placa, modelo, tipo_camion_id, es_rentado,
		        disponibilidad_id, nombre_disponibilidad, color_disponibilidad,
		        created_at, updated_at
		 FROM camion
		 WHERE eliminado = false`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var camiones []entities.Camion

	for rows.Next() {
		var c entities.Camion
		err := rows.Scan(
			&c.CamionID,
			&c.Placa,
			&c.Modelo,
			&c.TipoCamionID,
			&c.EsRentado,
			&c.DisponibilidadID,
			&c.NombreDisponibilidad,
			&c.ColorDisponibilidad,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		camiones = append(camiones, c)
	}

	return camiones, nil
}

func (pg *Postgres) Delete(id int32) error {
	cmd, err := pg.conn.Exec(context.Background(),
		`UPDATE camion SET eliminado = true WHERE camion_id = $1`, id)

	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("camion no encontrado")
	}
	return nil
}

func (pg *Postgres) GetByID(id int32) (*entities.Camion, error) {
	sql := `
	SELECT 
		camion_id,
		placa,
		modelo,
		tipo_camion_id,
		es_rentado,
		disponibilidad_id,
		nombre_disponibilidad,
		color_disponibilidad,
		created_at,
		updated_at
	FROM camion
	WHERE camion_id = $1 AND eliminado = false
	`

	var camion entities.Camion

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(
		&camion.CamionID,
		&camion.Placa,
		&camion.Modelo,
		&camion.TipoCamionID,
		&camion.EsRentado,
		&camion.DisponibilidadID,
		&camion.NombreDisponibilidad,
		&camion.ColorDisponibilidad,
		&camion.CreatedAt,
		&camion.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("camión no encontrado")
		}
		return nil, err
	}

	return &camion, nil
}


func (pg *Postgres) Update(id int32, camion *entities.Camion) (*entities.Camion, error) {
	camion.UpdatedAt = time.Now()
	sql := `
	UPDATE camion
	SET 
		placa = $1,
		modelo = $2,
		tipo_camion_id = $3,
		es_rentado = $4,
		disponibilidad_id = $5,
		nombre_disponibilidad = $6,
		color_disponibilidad = $7,
		updated_at = $8
	WHERE camion_id = $9 AND eliminado = false
	`

	cmdTag, err := pg.conn.Exec(
		context.Background(),
		sql,
		camion.Placa,
		camion.Modelo,
		camion.TipoCamionID,
		camion.EsRentado,
		camion.DisponibilidadID,
		camion.NombreDisponibilidad,
		camion.ColorDisponibilidad,
		camion.UpdatedAt, 
		id,
	)

	if err != nil {
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, errors.New("camión no encontrado")
	}

	return camion, nil
}


func (pg *Postgres) GetByPlaca(placa string) (*entities.Camion, error) {
	sql := `
	SELECT 
		camion_id,
		placa,
		modelo,
		tipo_camion_id,
		es_rentado,
		disponibilidad_id,
		nombre_disponibilidad,
		color_disponibilidad,
		created_at,
		updated_at
	FROM camion
	WHERE placa = $1 AND eliminado = false
	`

	var camion entities.Camion

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		placa,
	).Scan(
		&camion.CamionID,
		&camion.Placa,
		&camion.Modelo,
		&camion.TipoCamionID,
		&camion.EsRentado,
		&camion.DisponibilidadID,
		&camion.NombreDisponibilidad,
		&camion.ColorDisponibilidad,
		&camion.CreatedAt,
		&camion.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("camión no encontrado")
		}
		return nil, err
	}

	return &camion, nil
}

func (pg *Postgres) GetByModelo(modelo string) ([]entities.Camion, error) {
	sql := `
	SELECT 
		camion_id,
		placa,
		modelo,
		tipo_camion_id,
		es_rentado,
		disponibilidad_id,
		nombre_disponibilidad,
		color_disponibilidad,
		created_at,
		updated_at
	FROM camion
	WHERE modelo ILIKE '%' || $1 || '%' AND eliminado = false
	`

	rows, err := pg.conn.Query(context.Background(), sql, modelo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var camiones []entities.Camion

	for rows.Next() {
		var camion entities.Camion
		if err := rows.Scan(
			&camion.CamionID,
			&camion.Placa,
			&camion.Modelo,
			&camion.TipoCamionID,
			&camion.EsRentado,
			&camion.DisponibilidadID,
			&camion.NombreDisponibilidad,
			&camion.ColorDisponibilidad,
			&camion.CreatedAt,
			&camion.UpdatedAt,
		); err != nil {
			return nil, err
		}

		camiones = append(camiones, camion)
	}

	return camiones, nil
}
