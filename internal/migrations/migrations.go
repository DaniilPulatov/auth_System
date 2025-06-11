package migrations

import (
	internalErr "auth-service/internal/errors"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log"
)

func ApplyMigrations(filename, dsn string) error {
	migration, err := migrate.New("file://"+filename, dsn)
	if err != nil {
		log.Println("ApplyMigrations error:", err)
		return internalErr.ErrNewMigration
	}
	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Println("ApplyMigrations - Up() error:", err)
		return internalErr.ErrUpMigrations
	}

	log.Println("Migrations completed successfully")
	return nil
}
