package db

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*
var migrations embed.FS

func ConnectPostgresWithMigration(
	lg *log.Logger,
	postgresUser string,
	postgresPassword string,
	postgresDB string,
	dbHost string,
	sslmode string,
) (*sql.DB, error) {
	connstring := fmt.Sprintf(
		"user='%s' password='%s' dbname='%s' host='%s' sslmode='%s'",
		postgresUser,
		postgresPassword,
		postgresDB,
		dbHost,
		sslmode,
	)

	db, err := sql.Open("postgres", connstring)
	if err != nil {
		return &sql.DB{}, err
	}
	if err := db.Ping(); err != nil {
		return &sql.DB{}, err
	}

	srcDriver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return &sql.DB{}, err
	}
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return &sql.DB{}, err
	}
	migration, err := migrate.NewWithInstance(
		"iofs",
		srcDriver,
		"postgres",
		dbDriver,
	)
	if err != nil {
		return &sql.DB{}, err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return &sql.DB{}, err
	}

	return db, nil
}
