package configs

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
}
