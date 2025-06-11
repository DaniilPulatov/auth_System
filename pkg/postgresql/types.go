package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Ping(ctx context.Context) error
	Close()

	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type conn struct {
	*pgxpool.Pool
}

func (c conn) Ping(ctx context.Context) error {
	return c.Pool.Ping(ctx)
}

func (c conn) Close() {
	c.Pool.Close()
}

func (c conn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return c.Pool.QueryRow(ctx, sql, args...)
}

func (c conn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.Pool.Query(ctx, sql, args...)
}

func (c conn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.Pool.Exec(ctx, sql, arguments...)
}
