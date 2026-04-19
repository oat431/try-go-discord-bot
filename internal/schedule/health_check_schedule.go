package schedule

import (
	"oat431/try-go-discord-bot/internal/bot"
	"os"

	"github.com/bwmarrin/discordgo"
)

func HealthCheckSchedule() {
	healthCheckID := os.Getenv("HEALTH_CHECK_CHANNEL")
	if healthCheckID == "" {
		println("HEALTH_CHECK_CHANNEL is not set")
		return
	}
	s := bot.GetSession()

	_, err := s.ChannelMessageSendEmbed(
		healthCheckID,
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
		// Log the error instead of returning it, since this is a scheduled task
		// You can use a logging library or simply print the error
		println("Error sending health check message:", err.Error())
	}
}
