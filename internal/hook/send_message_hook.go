package hook

import (
	"oat431/try-go-discord-bot/internal/bot"
	"os"

	"github.com/gofiber/fiber/v3"
)

type messageRequest struct {
	Message   string `json:"message"`
	ChannelID string `json:"channel_id"`
}

func SendMessageHook(c fiber.Ctx) error {
	var payload messageRequest
	if err := c.Bind().Body(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	message := payload.Message
	if message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "message is required",
		})
	}

	channelID := payload.ChannelID
	if channelID == "" {
		channelID = os.Getenv("CHANNEL_ID")
	}
	if channelID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "channel_id is required",
		})
	}

	s := bot.GetSession()
	if s == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "discord session is not ready",
		})
	}

	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "message sent",
		"channel_id": channelID,
	})
}
