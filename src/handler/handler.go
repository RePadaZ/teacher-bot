package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"teacher-bot/src/text"
)

// DefaultHandler Хэндлера по умолчанию для если не сработали другие
func DefaultHandler(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	_, err := myBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text.NotHandler,
	})
	if err != nil {
		log.Default().Println(err)
	}
}

// Start Хэндлер для команды start с ответом
func Start(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	_, err := myBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text.StartBot,
	})
	if err != nil {
		log.Default().Println(err)
	}
}

// Help Хэндлер для команды help с ответом
func Help(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	_, err := myBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text.HelpBot,
	})
	if err != nil {
		log.Default().Println(err)
	}
}
