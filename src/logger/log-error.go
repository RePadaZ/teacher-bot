package logger

import (
	"errors"
	"github.com/go-telegram/bot"
	"log"
)

// LogError обрабатывает ошибки
func LogError(err error) {

	switch {
	case errors.Is(err, bot.ErrorForbidden):
		log.Println("Error 403: Access forbidden")
		break
	case errors.Is(err, bot.ErrorBadRequest):
		log.Println("Error 400: Bad request")
		break
	case errors.Is(err, bot.ErrorUnauthorized):
		log.Println("Error 401: Unauthorized")
		break
	case errors.Is(err, bot.ErrorNotFound):
		log.Println("Error 404: Not found")
		break
	case errors.Is(err, bot.ErrorConflict):
		log.Println("Error 409: Conflict")
		break
	case errors.Is(err, bot.ErrorTooManyRequests):
		log.Println("Error 429: Too many requests")
		break
	default:
		log.Println("Unknown error:", err)
	}

}
