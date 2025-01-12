package system

import (
	"github.com/joho/godotenv"
	"os"
)

func GetBotToken() string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	return os.Getenv("BOT_TOKEN")
}
