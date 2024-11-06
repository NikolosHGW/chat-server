package pg

import (
	"context"
	"database/sql"

	"github.com/NikolosHGW/chat-server/internal/client/db"
	"github.com/jmoiron/sqlx"
)

type pg struct {
	dbc *sqlx.DB
}

// NewDB - конструктор для постгрес-обёртки.
func NewDB(dbc *sqlx.DB) db.DB {
	return &pg{dbc: dbc}
}

func (p *pg) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return p.dbc.NamedExecContext(ctx, query, arg)
}

func (p *pg) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return p.dbc.SelectContext(ctx, dest, query, args...)
}

func (p *pg) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p.dbc.ExecContext(ctx, query, args...)
}

func (p *pg) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return p.dbc.QueryRowxContext(ctx, query, args...)
}

func (p *pg) PingContext(ctx context.Context) error {
	return p.dbc.PingContext(ctx)
}

func (p *pg) Close() error {
	return p.dbc.Close()
}
