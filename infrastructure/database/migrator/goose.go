package migrator

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq" // Cambia ao driver que uses
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type GooseClient struct {
	db      *sql.DB
	workdir string
	connStr string
}

func NewGooseClient(connStr string) (*GooseClient, error) {
	if connStr == "" {
		log.Error().Stack().Msg("No connection string provided")
		return nil, errors.New("no connection string provided")
	}

	// Abre DB
	db, err := sql.Open("postgres", connStr) // Asegúrate de cambiar "postgres" se usas outro motor
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to open database connection")
		return nil, err
	}

	// Verifica conexión
	if err := db.Ping(); err != nil {
		log.Error().Stack().Err(err).Msg("failed to ping database")
		return nil, err
	}

	// Directorio de migracións goose
	workdir := filepath.Join(".", "infrastructure", "database", "migrator", "migrations")
	if fi, err := os.Stat(workdir); err != nil || !fi.IsDir() {
		log.Error().Stack().Msgf("migrations dir not found: %s", workdir)
		return nil, fmt.Errorf("migrations dir not found: %s", workdir)
	}

	// Configura goose
	goose.SetBaseFS(os.DirFS(workdir))

	return &GooseClient{
		db:      db,
		workdir: workdir,
		connStr: connStr,
	}, nil
}

func (gc *GooseClient) ApplyMigrations(ctx context.Context) error {
	defer func() {
		if err := gc.db.Close(); err != nil {
			log.Error().Stack().Err(err).Msg("error closing database")
		}
	}()

	log.Info().Msg("Applying goose migrations...")

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(gc.db, "migrations"); err != nil {
		log.Error().Stack().Err(err).Msg("goose migration failed")
		return err
	}

	log.Info().Msg("Migrations applied successfully")
	return nil
}
