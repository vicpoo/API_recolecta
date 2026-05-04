package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/domain/entities"
	empleadoRepo "github.com/vicpoo/API_recolecta/src/empleado/infrastructure/repository"
)

type AdminSeedConfig struct {
	Nombre    string
	Apellidos string
	Mail      string
	Username  string
	Password  string
}

func SeedAdmin(ctx context.Context) error {
	cfg, err := loadAdminSeedConfig()
	if err != nil {
		return err
	}

	db := core.GetBD()
	defer core.ClosePool()

	seedCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := ensureAdminRole(seedCtx, db); err != nil {
		return err
	}

	repo := empleadoRepo.NewEmpleadoPostgresRepository(db)

	existingByMail, err := repo.FindByMail(seedCtx, cfg.Mail)
	if err != nil {
		return err
	}
	if existingByMail != nil {
		fmt.Printf("admin ya existe con mail %s (id=%d)\n", existingByMail.Mail, existingByMail.ID)
		return nil
	}

	existingByUsername, err := repo.FindByUsername(seedCtx, cfg.Username)
	if err != nil {
		return err
	}
	if existingByUsername != nil {
		fmt.Printf("admin ya existe con username %s (id=%d)\n", existingByUsername.Username, existingByUsername.ID)
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(cfg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	empleado := &entities.Empleado{
		Nombre:      cfg.Nombre,
		Apellidos:   cfg.Apellidos,
		Mail:        cfg.Mail,
		Username:    cfg.Username,
		Password:    string(hash),
		Desactivado: false,
		RolID:       core.ADMIN,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := repo.Create(seedCtx, empleado)
	if err != nil {
		return err
	}

	fmt.Printf("admin creado correctamente con id=%d\n", id)
	return nil
}

func loadAdminSeedConfig() (AdminSeedConfig, error) {
	cfg := AdminSeedConfig{
		Nombre:    strings.TrimSpace(os.Getenv("ADMIN_NOMBRE")),
		Apellidos: strings.TrimSpace(os.Getenv("ADMIN_APELLIDOS")),
		Mail:      strings.TrimSpace(strings.ToLower(os.Getenv("ADMIN_MAIL"))),
		Username:  strings.TrimSpace(strings.ToLower(os.Getenv("ADMIN_USERNAME"))),
		Password:  strings.TrimSpace(os.Getenv("ADMIN_PASSWORD")),
	}

	if cfg.Nombre == "" {
		cfg.Nombre = "Admin"
	}
	if cfg.Apellidos == "" {
		cfg.Apellidos = "Sistema"
	}

	var missing []string
	if cfg.Mail == "" {
		missing = append(missing, "ADMIN_MAIL")
	}
	if cfg.Username == "" {
		missing = append(missing, "ADMIN_USERNAME")
	}
	if cfg.Password == "" {
		missing = append(missing, "ADMIN_PASSWORD")
	}

	if len(missing) > 0 {
		return AdminSeedConfig{}, errors.New("faltan variables de entorno para el bootstrap: " + strings.Join(missing, ", "))
	}

	return cfg, nil
}

func ensureAdminRole(ctx context.Context, db *pgxpool.Pool) error {
	const q = `
		INSERT INTO rol (role_id, nombre, eliminado)
		VALUES ($1, $2, FALSE)
		ON CONFLICT (role_id)
		DO UPDATE SET
			nombre = EXCLUDED.nombre,
			eliminado = FALSE
	`

	_, err := db.Exec(ctx, q, core.ADMIN, "ADMIN")
	return err
}
