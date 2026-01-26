package adapters

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/entities"
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


func (postgres *Postgres) Save(tipoCamion *entities.TipoCamion) (*entities.TipoCamion, error) {
	tipoCamion.CreatedAt = time.Now()
	sql := `
	INSERT INTO tipo_camion
	(
		nombre,
		descripcion,
		created_at
	)
	VALUES ($1, $2, $3)
	RETURNING tipo_camion_id
	`

	err := postgres.conn.QueryRow(
		context.Background(),
		sql,
		tipoCamion.Nombre,
		tipoCamion.Descripcion,
		tipoCamion.CreatedAt, 
	).Scan(&tipoCamion.TipoCamionID)

	if err != nil {
		return nil, err
	}

	return tipoCamion, nil
}


func (posgres *Postgres) ListAll() ([]entities.TipoCamion, error) {
	var tipos []entities.TipoCamion
	sql := "SELECT * FROM tipo_camion"

	rows, err := posgres.conn.Query(context.Background(), sql)

	if err != nil {
		fmt.Printf("error:%s", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tipoCamion entities.TipoCamion

		if err = rows.Scan(&tipoCamion.TipoCamionID, &tipoCamion.Nombre, &tipoCamion.Descripcion, &tipoCamion.CreatedAt); err != nil {
			fmt.Printf("error; %s", err)
			return nil, err
		}

		tipos = append(tipos, tipoCamion)
	}

	if len(tipos) == 0 {
		return []entities.TipoCamion{}, nil
	}

	return tipos, nil
}

func (postgres *Postgres) GetByName(nombre string) (*entities.TipoCamion, error) {
	sql := `
		SELECT tipo_camion_id, nombre, descripcion, created_at
		FROM tipo_camion
		WHERE nombre = $1
		LIMIT 1
	`

	var tipoCamion entities.TipoCamion

	err := postgres.conn.QueryRow(
		context.Background(),
		sql,
		nombre,
	).Scan(
		&tipoCamion.TipoCamionID,
		&tipoCamion.Nombre,
		&tipoCamion.Descripcion,
		&tipoCamion.CreatedAt,
	)

	if err != nil {
		// no encontrado
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &tipoCamion, nil
}

func (postgres *Postgres) Delete(id int32) error {
	sql := "DELETE FROM tipo_camion WHERE tipo_camion_id = $1"

	result, err := postgres.conn.Exec(context.Background(), sql, id)
	
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("tipo de camión no encontrado")
	}

	return nil
}