package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
)

type Postgress struct {
	Pool *pgxpool.Pool
}

func New(dsn string) (*Postgress, error) {
	var pool *pgxpool.Pool
	var err error

	err = utils.DoWithtries(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, 5, 10*time.Second)
	return &Postgress{pool}, nil
}
