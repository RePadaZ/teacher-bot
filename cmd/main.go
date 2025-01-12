package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"teacher-bot/pkg/system"
	"teacher-bot/src/filter"
	"teacher-bot/src/handler"
	"teacher-bot/src/text"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Получение токена
	token := system.GetBotToken()

	// Загрузка хэндлера по умолчанию для если не сработали другие
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	// Старт бота
	myBot, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	// Регистрируем наши хэндлеры
	myBot.RegisterHandlerMatchFunc(filter.IsStar, handler.Start)

	myBot.Start(ctx)
}

// defaultHandler Хэндлера по умолчанию для если не сработали другие
func defaultHandler(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	_, err := myBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text.NotHandler,
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
}
