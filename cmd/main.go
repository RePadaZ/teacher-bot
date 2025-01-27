package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"teacher-bot/pkg/system"
	"teacher-bot/src/initBot"
)

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
	myBot, err := initBot.InitializeBot(token)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	// Регистрация хэндлеров
	initBot.RegisterHandlers(myBot)

	log.Println("Bot is now running. Press CTRL-F2 to exit.")

	// Запуск бота
	myBot.Start(ctx)

	// Логирование при завершении работы
	log.Println("Bot stopped gracefully.")
}
