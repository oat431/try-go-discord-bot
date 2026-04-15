package behavior

import (
	"fmt"
	"oat431/try-go-discord-bot/internal/command"
	"oat431/try-go-discord-bot/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func StartDiscordBot() {
	config.LoadEnvConfig()

	Token := os.Getenv("TOKEN")
	if Token == "" {
		fmt.Println("DISCORD_BOT_TOKEN environment variable is not set")
		return
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// register command here
	dg.AddHandler(command.PingCommand)
	dg.AddHandler(command.HandleSlashCommand)
	// end of command registration

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	err = command.CleanupGlobalSlashCommands(dg)
	if err != nil {
		fmt.Println("error cleaning up global slash commands,", err)
		return
	}

	err = command.RegisterSlashCommands(dg)
	if err != nil {
		fmt.Println("error registering slash commands,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
