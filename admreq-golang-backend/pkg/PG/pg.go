package pg

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
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

	ct, err := pool.Exec(context.Background(), "SELECT * FROM requsers")
	if err != nil {
		return nil, fmt.Errorf("error while checking users in base: %v", err)
	}
	if ct.RowsAffected() == 0 {
		var dep_id string
		err = pool.QueryRow(context.Background(), "INSERT INTO departments (department_name) VALUES ('admin')").Scan(dep_id)
		if err != nil {
			return nil, fmt.Errorf("error while adding default department: %v", err)
		}
		_, err = pool.Exec(context.Background(), "INSERT INTO requsers (firstname, lastname, surname, department, user_role, user_login, user_pass) VALUES ('admin', 'admin', 'admin', $1, 'admin', 'admin', $2)", dep_id, utils.HashPassword("admin"))
		if err != nil {
			return nil, fmt.Errorf("error while adding default user: %v", err)
		}
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
