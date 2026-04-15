package utils

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

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

	fmt.Println("finished cleaning old command")

	return nil
}

func RegisterSlashCommands(s *discordgo.Session, commands []*discordgo.ApplicationCommand) error {
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

	fmt.Println("finished registering slash commands")

	return nil
}
