package main

import (
	"context"

	"github.com/E7ast1c/Cupbearer/configs"
	"github.com/E7ast1c/Cupbearer/internal/database"
	"github.com/E7ast1c/Cupbearer/internal/telegram"
)

func main() {
	ctx := context.Background()
	configs.InitEnv()
	database.NewPGConn(ctx)
	telegram.TGConnect()
}
