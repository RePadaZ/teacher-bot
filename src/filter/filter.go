package filter

import (
	"github.com/go-telegram/bot/models"
	"strings"
)

// IsStart Фильтр для команды старт
func IsStart(update *models.Update) bool {
	return strings.ToLower(update.Message.Text) == "/start"
}
