package repository

import (
	"context"
	"database/sql"
)

type contextExecutor interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func (c *carRepository) exec(isWrite bool) contextExecutor {
	if c.tx != nil {
		return c.tx
	}
	if isWrite {
		return c.dbWrite
	}
	return c.dbRead
}

func (e *engineRepository) exec(isWrite bool) contextExecutor {
	if e.tx != nil {
		return e.tx
	}
	if isWrite {
		return e.dbWrite
	}
	return e.dbRead
}
