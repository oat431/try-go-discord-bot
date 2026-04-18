package bot

import (
	"fmt"
	"oat431/try-go-discord-bot/internal/command"
	"oat431/try-go-discord-bot/pkg/utils"
	"os"

	"github.com/bwmarrin/discordgo"
)

func StartDiscordBot() (*discordgo.Session, error) {
	Token := os.Getenv("TOKEN")
	if Token == "" {
		return nil, fmt.Errorf("TOKEN environment variable is not set")
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	// Register the slash command handler
	dg.AddHandler(command.HandleSlashCommand)

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %w", err)
	}

	err = utils.CleanupGlobalSlashCommands(dg)
	if err != nil {
		return nil, fmt.Errorf("error cleaning up global slash commands: %w", err)
	}

	err = utils.RegisterSlashCommands(dg, command.Commands)
	if err != nil {
		return nil, fmt.Errorf("error registering slash commands: %w", err)
	}

	setSession(dg)
	fmt.Println("Bot is now running.")

	return dg, nil
}
