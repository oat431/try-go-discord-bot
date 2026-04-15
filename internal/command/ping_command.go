package command

import (
	"github.com/bwmarrin/discordgo"
)

func PingSlashInfo() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with Pong!",
	}
}

func PongSlashInfo() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "pong",
		Description: "Replies with Ping!",
	}
}

func PingSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}

func PongSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ping!",
		},
	})
}
