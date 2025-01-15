package main

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"os"
	"os/signal"
	"teacher-bot/pkg/system"
	"teacher-bot/src/filter"
	"teacher-bot/src/handler"
)

// initializeBot создает и настраивает бота
func initializeBot(token string) (*bot.Bot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler), // Хэндлер по умолчанию
		bot.WithSkipGetMe(),                            // Пропуск вызова getMe при запуске
		bot.WithAllowedUpdates(bot.AllowedUpdates{
			"message",
		}),
		bot.WithInitialOffset(int64(-2)), // Установка начального оффсета
	}

	return bot.New(token, opts...)
}

// registerHandlers регистрирует хэндлеры для бота
func registerHandlers(myBot *bot.Bot) {
	myBot.RegisterHandlerMatchFunc(filter.IsStart, handler.Start)
	myBot.RegisterHandlerMatchFunc(filter.IsHelp, handler.Help)
}

func main() {
	// Настройка контекста с обработкой сигналов завершения
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Получение токена и проверка на его наличие
	token := system.GetBotToken()
	if token == "" {
		log.Fatal("Bot token is missing. Please provide a valid token.")
	}

	// Инициализация бота
	myBot, err := initializeBot(token)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	// Регистрация хэндлеров
	registerHandlers(myBot)

	log.Println("Bot is now running. Press CTRL-F2 to exit.")

	// Запуск бота
	myBot.Start(ctx)

	// Логирование при завершении работы
	log.Println("Bot stopped gracefully.")
}
