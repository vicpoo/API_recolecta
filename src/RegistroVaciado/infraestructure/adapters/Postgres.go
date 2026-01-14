package adapters

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vicpoo/API_recolecta/src/RegistroVaciado/domain/entities"
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
// ================== CREATE ==================
//
func (pg *Postgres) Save(r *entities.RegistroVaciado) (*entities.RegistroVaciado, error) {
	sql := `
	INSERT INTO registro_vaciado (relleno_id, ruta_camion_id, hora)
	VALUES ($1, $2, $3)
	RETURNING vaciado_id
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		r.RellenoID,
		r.RutaCamionID,
		r.Hora,
	).Scan(&r.VaciadoID)

	if err != nil {
		return nil, err
	}

	return r, nil
}

//
// ================== GET BY ID ==================
//
func (pg *Postgres) GetByID(id int32) (*entities.RegistroVaciado, error) {
	sql := `
	SELECT vaciado_id, relleno_id, ruta_camion_id, hora
	FROM registro_vaciado
	WHERE vaciado_id = $1
	`

	return pg.fetchOne(sql, id)
}

//
// ================== LIST ALL ==================
//
func (pg *Postgres) ListAll() ([]entities.RegistroVaciado, error) {
	sql := `
	SELECT vaciado_id, relleno_id, ruta_camion_id, hora
	FROM registro_vaciado
	ORDER BY vaciado_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entities.RegistroVaciado

	for rows.Next() {
		var r entities.RegistroVaciado
		err := rows.Scan(
			&r.VaciadoID,
			&r.RellenoID,
			&r.RutaCamionID,
			&r.Hora,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}

	return list, nil
}

//
// ================== DELETE ==================
//
func (pg *Postgres) Delete(id int32) error {
	_, err := pg.conn.Exec(
		context.Background(),
		`DELETE FROM registro_vaciado WHERE vaciado_id=$1`,
		id,
	)
	return err
}

//
// ================== GET BY RELLENO ==================
//
func (pg *Postgres) GetByRellenoID(rellenoID int32) ([]entities.RegistroVaciado, error) {
	sql := `
	SELECT vaciado_id, relleno_id, ruta_camion_id, hora
	FROM registro_vaciado
	WHERE relleno_id = $1
	ORDER BY hora DESC
	`

	return pg.fetchMany(sql, rellenoID)
}

//
// ================== GET BY RUTA CAMION ==================
//
func (pg *Postgres) GetByRutaCamionID(rutaCamionID int32) ([]entities.RegistroVaciado, error) {
	sql := `
	SELECT vaciado_id, relleno_id, ruta_camion_id, hora
	FROM registro_vaciado
	WHERE ruta_camion_id = $1
	ORDER BY hora DESC
	`

	return pg.fetchMany(sql, rutaCamionID)
}

//
// ================== EXISTS ==================
//
func (pg *Postgres) ExistsByID(id int32) (bool, error) {
	sql := `SELECT EXISTS(SELECT 1 FROM registro_vaciado WHERE vaciado_id=$1)`
	var exists bool

	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(&exists)
	return exists, err
}

//
// ================== HELPERS ==================
//
func (pg *Postgres) fetchOne(sql string, param any) (*entities.RegistroVaciado, error) {
	var r entities.RegistroVaciado

	err := pg.conn.QueryRow(context.Background(), sql, param).Scan(
		&r.VaciadoID,
		&r.RellenoID,
		&r.RutaCamionID,
		&r.Hora,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("registro_vaciado no encontrado")
		}
		return nil, err
	}

	return &r, nil
}

func (pg *Postgres) fetchMany(sql string, param any) ([]entities.RegistroVaciado, error) {
	rows, err := pg.conn.Query(context.Background(), sql, param)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entities.RegistroVaciado

	for rows.Next() {
		var r entities.RegistroVaciado
		err := rows.Scan(
			&r.VaciadoID,
			&r.RellenoID,
			&r.RutaCamionID,
			&r.Hora,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}

	return list, nil
}
