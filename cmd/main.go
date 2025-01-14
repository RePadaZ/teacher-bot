package main

import (
	"context"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
	"teacher-bot/pkg/system"
	"teacher-bot/src/filter"
	"teacher-bot/src/handler"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Получение токена
	token := system.GetBotToken()

	// Загрузка хэндлера по умолчанию для если не сработали другие
	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler),
		bot.WithSkipGetMe(),
		bot.WithAllowedUpdates(bot.AllowedUpdates{
			"message",
		}),
		bot.WithInitialOffset(int64(-2)),
	}

	// Старт бота
	myBot, err := bot.New(token, opts...)
	if err != nil {
		panic(err.Error())
	}

	// Регистрируем наши хэндлеры
	myBot.RegisterHandlerMatchFunc(filter.IsStart, handler.Start)
	myBot.RegisterHandlerMatchFunc(filter.IsHelp, handler.Help)

	myBot.Start(ctx)
}
