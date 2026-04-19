package main

import (
	"fmt"
	"oat431/try-go-discord-bot/internal/api"
	"oat431/try-go-discord-bot/internal/bot"
	"oat431/try-go-discord-bot/internal/config"
	"oat431/try-go-discord-bot/internal/schedule"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadEnvConfig()

	dg, err := bot.StartDiscordBot()
	if err != nil {
		fmt.Println(err)
		return
	}

	go api.StartAPI()

	stopSchedule := make(chan struct{})
	go schedule.StartHealthCheckTicker(stopSchedule, time.Minute)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	close(stopSchedule)

	if dg != nil {
		dg.Close()
	}
}
