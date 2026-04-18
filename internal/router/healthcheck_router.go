package router

import (
	"oat431/try-go-discord-bot/internal/hook"

	"github.com/gofiber/fiber/v3"
)

func RegisterHealthCheckRoute(router fiber.Router) {
	router.Get("/health", hook.HealthCheckHook)
}
