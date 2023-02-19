package database

import (
	"context"
	"fmt"
	"os"

	Config "github.com/E7ast1c/Cupbearer/config"
	"github.com/jackc/pgx/v4"
)

type PGConn struct {
	Ctx  context.Context
	Conn *pgx.Conn
}

func NewPGConn(ctx context.Context, pgConfig Config.PgConfig) *PGConn {
	return &PGConn{Ctx: ctx, Conn: connect(ctx, pgConfig)}
}

func connect(ctx context.Context, pgConfig Config.PgConfig) *pgx.Conn {
	conn, err := pgx.Connect(ctx, pgConfig.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
