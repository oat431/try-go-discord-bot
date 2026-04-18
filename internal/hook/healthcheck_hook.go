package hook

import (
	"oat431/try-go-discord-bot/internal/bot"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v3"
)

func HealthCheckHook(c fiber.Ctx) error {
	channelID := os.Getenv("CHANNEL_ID")
	s := bot.GetSession()

	_, err := s.ChannelMessageSendEmbed(
		channelID,
		&discordgo.MessageEmbed{
			Title:       "Health Check",
			Description: "This is a health check message from the API.",
			Color:       0x00ff00,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Status",
					Value:  "Healthy",
					Inline: true,
				},
			},
			Timestamp: "",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Health check successful",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.imgur.com/4M34hi2.png",
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.imgur.com/4M34hi2.png",
			},
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Health check successful",
		"channel_id": channelID,
	})
}
