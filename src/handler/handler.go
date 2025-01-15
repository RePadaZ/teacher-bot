package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"teacher-bot/src/logger"
	"teacher-bot/src/text"
)

// DefaultHandler обрабатывает неизвестные команды
func DefaultHandler(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	if err := sendMessage(ctx, myBot, update.Message.Chat.ID, text.NotHandler); err != nil {
		logger.LogError(err)
	}
}

// Start обрабатывает команду /start
func Start(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	if err := sendMessage(ctx, myBot, update.Message.Chat.ID, text.StartBot); err != nil {
		logger.LogError(err)
	}
}

// Help обрабатывает команду /help
func Help(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	if err := sendMessage(ctx, myBot, update.Message.Chat.ID, text.HelpBot); err != nil {
		logger.LogError(err)
	}
}

// sendMessage отправляет сообщение пользователю
func sendMessage(ctx context.Context, myBot *bot.Bot, chatID int64, message string) error {
	_, err := myBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   message,
	})
	return err
}
