package pg

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PG struct {
	*pgxpool.Pool
}

func NewDB(dsn, mp string) (*PG, error) {
	var pool *pgxpool.Pool
	var err error

	err = doMigration(dsn, mp)
	if err != nil {
		return nil, fmt.Errorf("error while do migration in database: %v", err)
	}

	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("error while creating pgx pool: %v", err)
	}
	return &PG{pool}, nil
}

func doMigration(dsn, mp string) error {
	m, err := migrate.New(mp, dsn)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
