package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type PGConn struct {
	Ctx  context.Context
	Conn *pgx.Conn
}

func NewPGConn(ctx context.Context) *PGConn {
	return &PGConn{Ctx: ctx, Conn: connect(ctx)}
}

func connect(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
