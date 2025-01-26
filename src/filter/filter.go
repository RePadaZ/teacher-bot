package filter

import (
	"github.com/go-telegram/bot/models"
	"strings"
)

// IsStart Фильтр для команды старт
func IsStart(update *models.Update) bool {
	return update.Message != nil && strings.ToLower(update.Message.Text) == "/start"
}

// IsHelp Фильтр для команды хелп
func IsHelp(update *models.Update) bool {
	return update.Message != nil && strings.ToLower(update.Message.Text) == "/help"
}
