package postgresql

import (
	pkgErrors "auth-service/pkg/errors"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewPostgresDB(dsn string) (Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Println(err)
		return nil, pkgErrors.ErrParseConfig
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Println(err)
		return nil, pkgErrors.ErrNewWithConfig
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Println(err)
		return nil, pkgErrors.ErrPingConnection
	}
	log.Println("Database connection established successfully")
	return conn{pool}, nil
}
