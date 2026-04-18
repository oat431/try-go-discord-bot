package router

import (
	"oat431/try-go-discord-bot/internal/hook"

	"github.com/gofiber/fiber/v3"
)

func RegisterSendMessageRoute(router fiber.Router) {
	router.Post("/message", hook.SendMessageHook)
}
