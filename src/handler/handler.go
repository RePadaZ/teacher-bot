package handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"teacher-bot/src/text"
)

// Start Хэндлер для команды старт с ответом
func Start(ctx context.Context, myBot *bot.Bot, update *models.Update) {
	_, err := myBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text.StartBot,
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
}
