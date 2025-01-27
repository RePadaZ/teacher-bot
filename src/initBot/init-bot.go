package initBot

import (
	"github.com/go-telegram/bot"
	"teacher-bot/src/filter"
	"teacher-bot/src/handler"
)

// InitializeBot создает и настраивает бота
func InitializeBot(token string) (*bot.Bot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler), // Хэндлер по умолчанию
		bot.WithSkipGetMe(),                            // Пропуск вызова getMe при запуске
		bot.WithInitialOffset(int64(-2)),               // Установка начального оффсета
	}

	return bot.New(token, opts...)
}

// RegisterHandlers регистрирует хэндлеры для бота
func RegisterHandlers(myBot *bot.Bot) {
	myBot.RegisterHandlerMatchFunc(filter.IsStart, handler.Start)
	myBot.RegisterHandlerMatchFunc(filter.IsHelp, handler.Help)
}
