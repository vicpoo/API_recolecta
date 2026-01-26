package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/domain/entities"
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
// SAVE
//
func (pg *Postgres) Save(p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error) {
	p.CreatedAt = time.Now()
	sql := `
	INSERT INTO punto_recoleccion
	(
		ruta_id,
		cp,
		eliminado,
		created_at
	)
	VALUES ($1, $2, false, $3)
	RETURNING punto_id
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		p.RutaID,
		p.CP,
		p.CreatedAt, // 👈 tú lo insertas
	).Scan(&p.PuntoID)

	if err != nil {
		return nil, err
	}

	return p, nil
}


//
// UPDATE
//
func (pg *Postgres) Update(id int32, p *entities.PuntoRecoleccion) (*entities.PuntoRecoleccion, error) {
	sql := `
	UPDATE punto_recoleccion
	SET
		ruta_id = $1,
		cp = $2
	WHERE punto_id = $3 AND eliminado = false
	`

	ct, err := pg.conn.Exec(
		context.Background(),
		sql,
		p.RutaID,
		p.CP,
		id,
	)

	if err != nil {
		return nil, err
	}

	if ct.RowsAffected() == 0 {
		return nil, errors.New("punto de recolección no encontrado")
	}

	return p, nil
}

//
// GET ALL
//
func (pg *Postgres) ListAll() ([]entities.PuntoRecoleccion, error) {
	sql := `
	SELECT punto_id, ruta_id, cp, eliminado
	FROM punto_recoleccion
	WHERE eliminado = false
	ORDER BY punto_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var puntos []entities.PuntoRecoleccion

	for rows.Next() {
		var p entities.PuntoRecoleccion
		err := rows.Scan(&p.PuntoID, &p.RutaID, &p.CP, &p.Eliminado)
		if err != nil {
			return nil, err
		}
		puntos = append(puntos, p)
	}

	return puntos, nil
}

//
// GET BY ID
//
func (pg *Postgres) GetById(id int32) (*entities.PuntoRecoleccion, error) {
	var p entities.PuntoRecoleccion

	sql := `
	SELECT punto_id, ruta_id, cp, eliminado
	FROM punto_recoleccion
	WHERE punto_id = $1 AND eliminado = false
	`

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(
		&p.PuntoID,
		&p.RutaID,
		&p.CP,
		&p.Eliminado,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("punto de recolección no encontrado")
		}
		return nil, err
	}

	return &p, nil
}

//
// GET BY RUTA
//
func (pg *Postgres) GetByRuta(rutaId int32) ([]entities.PuntoRecoleccion, error) {
	sql := `
	SELECT punto_id, ruta_id, cp, eliminado
	FROM punto_recoleccion
	WHERE ruta_id = $1 AND eliminado = false
	ORDER BY punto_id
	`

	rows, err := pg.conn.Query(context.Background(), sql, rutaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var puntos []entities.PuntoRecoleccion

	for rows.Next() {
		var p entities.PuntoRecoleccion
		err := rows.Scan(&p.PuntoID, &p.RutaID, &p.CP, &p.Eliminado)
		if err != nil {
			return nil, err
		}
		puntos = append(puntos, p)
	}

	return puntos, nil
}

//
// DELETE (Soft delete)
//
func (pg *Postgres) Delete(id int32) error {
	sql := `
	UPDATE punto_recoleccion
	SET eliminado = true
	WHERE punto_id = $1
	`

	ct, err := pg.conn.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return errors.New("punto de recolección no encontrado")
	}

	return nil
}
