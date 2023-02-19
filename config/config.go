package Config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func BuildAppConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file", err)
	}

	return AppConfig{
		Tg: TgConfig{
			Token:   GetEnvStr("TG_BOT_TOKEN"),
			BotName: GetEnvStr("TG_BOT_NAME"),
			ChatId:  GetEnvStr("TG_CHAT_ID"),
		},
		Pg: PgConfig{
			DbUrl: GetEnvStr("DB_URL"),
		},
	}
}

type AppConfig struct {
	Tg TgConfig
	Pg PgConfig
}

type TgConfig struct {
	Token   string
	BotName string
	ChatId  string
}

type PgConfig struct {
	DbUrl string
}

func GetEnvStr(name string) string {
	env := os.Getenv(name)
	if env == "" {
		logrus.Fatalf("read empty env: %s", name)
	}
	return env
}
