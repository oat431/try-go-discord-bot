package main

import (
	"fmt"
	"oat431/try-go-discord-bot/internal/api"
	"oat431/try-go-discord-bot/internal/bot"
	"oat431/try-go-discord-bot/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadEnvConfig()

	dg, err := bot.StartDiscordBot()
	if err != nil {
		fmt.Println(err)
		return
	}

	go api.StartAPI()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if dg != nil {
		dg.Close()
	}
}
