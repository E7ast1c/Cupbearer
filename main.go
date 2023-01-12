package main

import (
	"context"

	"github.com/E7ast1c/Cupbearer/configs"
	"github.com/E7ast1c/Cupbearer/internal/cron"
	"github.com/E7ast1c/Cupbearer/internal/database"
	job "github.com/E7ast1c/Cupbearer/internal/job"
	"github.com/E7ast1c/Cupbearer/internal/telegram"
)

func main() {
	ctx := context.Background()
	configs.InitEnv()
	dbConn := database.NewPGConn(ctx)
	telegram.TGConnect()

	cron.RunCron(job.NewHandler(dbConn.Ctx, dbConn.Conn))
}
