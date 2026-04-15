package main

import (
	"oat431/try-go-discord-bot/internal/bot"
	"oat431/try-go-discord-bot/internal/config"
)

func main() {
	config.LoadEnvConfig()
	bot.StartDiscordBot()
}
