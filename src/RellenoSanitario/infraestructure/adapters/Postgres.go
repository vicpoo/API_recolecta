package adapters

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/entities"
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
func (pg *Postgres) Save(relleno *entities.RellenoSanitario) (*entities.RellenoSanitario, error) {
	sql := `
	INSERT INTO relleno_sanitario (
		nombre,
		direccion,
		es_rentado,
		capacidad_toneladas
	)
	VALUES ($1, $2, $3, $4)
	RETURNING
		relleno_id,
		eliminado
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		relleno.Nombre,
		relleno.Direccion,
		relleno.EsRentado,
		relleno.CapacidadToneladas,
	).Scan(
		&relleno.RellenoID,
		&relleno.Eliminado,
	)

	if err != nil {
		return nil, err
	}

	return relleno, nil
}

//
// GET BY ID
//
func (pg *Postgres) GetByID(id int) (*entities.RellenoSanitario, error) {
	var relleno entities.RellenoSanitario

	sql := `
	SELECT
		relleno_id,
		nombre,
		direccion,
		es_rentado,
		eliminado,
		capacidad_toneladas
	FROM relleno_sanitario
	WHERE relleno_id = $1
	  AND eliminado = false
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		id,
	).Scan(
		&relleno.RellenoID,
		&relleno.Nombre,
		&relleno.Direccion,
		&relleno.EsRentado,
		&relleno.Eliminado,
		&relleno.CapacidadToneladas,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("relleno_sanitario no encontrado")
		}
		return nil, err
	}

	return &relleno, nil
}

//
// LIST ALL
//
func (pg *Postgres) ListAll() ([]entities.RellenoSanitario, error) {
	sql := `
	SELECT
		relleno_id,
		nombre,
		direccion,
		es_rentado,
		eliminado,
		capacidad_toneladas
	FROM relleno_sanitario
	WHERE eliminado = false
	ORDER BY relleno_id DESC
	`

	rows, err := pg.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rellenos []entities.RellenoSanitario

	for rows.Next() {
		var r entities.RellenoSanitario
		err := rows.Scan(
			&r.RellenoID,
			&r.Nombre,
			&r.Direccion,
			&r.EsRentado,
			&r.Eliminado,
			&r.CapacidadToneladas,
		)
		if err != nil {
			return nil, err
		}
		rellenos = append(rellenos, r)
	}

	return rellenos, nil
}

//
// UPDATE
//
func (pg *Postgres) Update(id int32, relleno *entities.RellenoSanitario) (*entities.RellenoSanitario, error) {
	sql := `
	UPDATE relleno_sanitario
	SET
		nombre = $1,
		direccion = $2,
		es_rentado = $3,
		capacidad_toneladas = $4
	WHERE relleno_id = $5
	  AND eliminado = false
	RETURNING eliminado
	`

	err := pg.conn.QueryRow(
		context.Background(),
		sql,
		relleno.Nombre,
		relleno.Direccion,
		relleno.EsRentado,
		relleno.CapacidadToneladas,
		id,
	).Scan(&relleno.Eliminado)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("relleno_sanitario no encontrado")
		}
		return nil, err
	}

	return relleno, nil
}

//
// DELETE (lógico)
//
func (pg *Postgres) Delete(id int) error {
	sql := `
	UPDATE relleno_sanitario
	SET eliminado = true
	WHERE relleno_id = $1
	`

	_, err := pg.conn.Exec(context.Background(), sql, id)
	return err
}

//
// GET BY NOMBRE
//
func (pg *Postgres) GetByNombre(nombre string) ([]entities.RellenoSanitario, error) {
	sql := `
	SELECT
		relleno_id,
		nombre,
		direccion,
		es_rentado,
		eliminado,
		capacidad_toneladas
	FROM relleno_sanitario
	WHERE LOWER(nombre) LIKE LOWER($1)
	  AND eliminado = false
	`

	rows, err := pg.conn.Query(context.Background(), sql, "%"+nombre+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rellenos []entities.RellenoSanitario

	for rows.Next() {
		var r entities.RellenoSanitario
		err := rows.Scan(
			&r.RellenoID,
			&r.Nombre,
			&r.Direccion,
			&r.EsRentado,
			&r.Eliminado,
			&r.CapacidadToneladas,
		)
		if err != nil {
			return nil, err
		}
		rellenos = append(rellenos, r)
	}

	return rellenos, nil
}

//
// EXISTS BY ID
//
func (pg *Postgres) ExistsByID(id int) (bool, error) {
	sql := `
	SELECT EXISTS (
		SELECT 1
		FROM relleno_sanitario
		WHERE relleno_id = $1
		  AND eliminado = false
	)
	`

	var exists bool
	err := pg.conn.QueryRow(context.Background(), sql, id).Scan(&exists)
	return exists, err
}
