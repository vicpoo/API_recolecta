package adapters

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/TipoCamion/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func NewPosgres() *Postgres {
	conn, _ := core.ConnectPostgres()
	return &Postgres{
		conn: conn,
	}
}


func (posgres *Postgres) Save(tipoCamion *entities.TipoCamion) (*entities.TipoCamion, error) {
	sql := "INSERT INTO tipo_camion(nombre, descripcion) VALUES ($1, $2) RETURNING tipo_camion_id, created_at"

	err := posgres.conn.QueryRow(context.Background(), sql, tipoCamion.Nombre, tipoCamion.Descripcion).Scan(&tipoCamion.TipoCamionID, &tipoCamion.CreatedAt)

	if err != nil {
		log.Fatal(err)
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

	result, err := postgres.conn.Exec(context.Background(), sql, id); 

	if err != nil {
		fmt.Printf("error: %s", err)
		return err
	}

	if result.RowsAffected() == 0 {
		log.Fatal(errors.New("error no se encontro el tipo de camion"))
	}

	return nil
}