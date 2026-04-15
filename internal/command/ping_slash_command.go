package command

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with Pong!",
		},
		{
			Name:        "pong",
			Description: "Replies with Ping!",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
		"pong": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Ping!",
				},
			})
		},
	}
)

func HandleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}

func getApplicationID(s *discordgo.Session) (string, error) {
	applicationID := s.State.User.ID
	if applicationID == "" {
		currentUser, err := s.User("@me")
		if err != nil {
			return "", fmt.Errorf("failed to resolve bot user: %w", err)
		}
		applicationID = currentUser.ID
	}

	return applicationID, nil
}

func CleanupGlobalSlashCommands(s *discordgo.Session) error {
	applicationID, err := getApplicationID(s)
	if err != nil {
		return err
	}

	globalCommands, err := s.ApplicationCommands(applicationID, "")
	if err != nil {
		return fmt.Errorf("failed to fetch global slash commands: %w", err)
	}

	for _, globalCommand := range globalCommands {
		err = s.ApplicationCommandDelete(applicationID, "", globalCommand.ID)
		if err != nil {
			return fmt.Errorf("failed to delete global slash command %q: %w", globalCommand.Name, err)
		}
	}

	return nil
}

func RegisterSlashCommands(s *discordgo.Session) error {
	guildID := os.Getenv("GUILD_ID")
	if guildID == "" {
		return fmt.Errorf("GUILD_ID environment variable is not set")
	}

	applicationID, err := getApplicationID(s)
	if err != nil {
		return err
	}

	_, err = s.ApplicationCommandBulkOverwrite(applicationID, guildID, commands)
	if err != nil {
		return fmt.Errorf("failed to overwrite guild slash commands: %w", err)
	}

	return nil
}
