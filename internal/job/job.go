package job

import (
	"context"

	Config "github.com/E7ast1c/Cupbearer/config"
	"github.com/jackc/pgx/v4"
)

type Job interface {
	Do()
	Cancel()
}

type Handler struct {
	Ctx       context.Context
	Conn      *pgx.Conn
	AppConfig Config.AppConfig
}

func NewHandler(ctx context.Context, conn *pgx.Conn, appConfig Config.AppConfig) *Handler {
	return &Handler{Ctx: ctx, Conn: conn, AppConfig: appConfig}
}
