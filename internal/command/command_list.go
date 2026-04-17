package command

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		PingSlashInfo(),
		PongSlashInfo(),
		FactorialSlashInfo(),
		IsPrimeSlashInfo(),
		WelcomSlashInfo(),
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":      PingSlashCommand,
		"pong":      PongSlashCommand,
		"factorial": FactorialSlashCommand,
		"isprime":   IsPrimeSlashCommand,
		"welcome":   WelcomSlashCommand,
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
