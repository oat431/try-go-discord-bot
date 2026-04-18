package command

import (
	"github.com/bwmarrin/discordgo"
)

func WelcomSlashInfo() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "welcome",
		Description: "Send a welcome embed message.",
	}
}

func WelcomSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.Member.User
	if user == nil {
		user = i.User
	}
	mention := "there"
	if user != nil {
		mention = user.Mention()
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Welcome!",
					Description: "Glad you are here. Feel free to check the rules and say hi!",
					Color:       0x57F287,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Member",
							Value:  mention,
							Inline: true,
						},
						{
							Name:   "Getting Started",
							Value:  "1) Read #rules\n2) Pick roles in #roles\n3) Introduce yourself in #introductions",
							Inline: false,
						},
					},
					Footer: &discordgo.MessageEmbedFooter{
						Text: "If you need help, ping a moderator.",
					},
				},
			},
		},
	})
}
