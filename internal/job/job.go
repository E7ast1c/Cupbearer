package job

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Job interface {
	Do()
	Cancel()
}

type Handler struct {
	Ctx  context.Context
	Conn *pgx.Conn
}

func NewHandler(ctx context.Context, conn *pgx.Conn) *Handler {
	return &Handler{Ctx: ctx, Conn: conn}
}
