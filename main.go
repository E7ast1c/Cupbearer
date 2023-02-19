package main

import (
	"context"

	Config "github.com/E7ast1c/Cupbearer/config"
	"github.com/E7ast1c/Cupbearer/internal/database"
	"github.com/E7ast1c/Cupbearer/internal/job"
)

func main() {
	ctx := context.Background()

	appConfig := Config.BuildAppConfig()
	dbConn := database.NewPGConn(ctx, appConfig.Pg)
	// fiat_rate.CbrRate{TgConfig: appConfig.Tg}.Do()

	handler := job.NewHandler(dbConn.Ctx, dbConn.Conn, appConfig)
	handler.RunCron(ctx)
}
