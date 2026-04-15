package main

import (
	"oat431/try-go-discord-bot/internal/behavior"
	"oat431/try-go-discord-bot/internal/config"
)

func main() {
	config.LoadEnvConfig()
	behavior.StartDiscordBot()
}
