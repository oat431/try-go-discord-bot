package command

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with Pong!",
		},
		{
			Name:        "pong",
			Description: "Replies with Ping!",
		},
		{
			Name:        "factorial",
			Description: "Calculates the factorial of a number.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number",
					Description: "The number to calculate factorial for.",
					Required:    true,
				},
			},
		},
		{
			Name:        "isprime",
			Description: "Checks if a number is prime.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number",
					Description: "The number to check for primality.",
					Required:    true,
				},
			},
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":      PingSlashCommand,
		"pong":      PongSlashCommand,
		"factorial": FactorialSlashCommand,
		"isprime":   IsPrimeSlashCommand,
	}
)

func HandleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if handler, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}
