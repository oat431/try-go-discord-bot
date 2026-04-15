package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("message:", m.Content)
	fmt.Println("authorID:", m.Author.ID)
	fmt.Println("stateUSER:", s.State.User.ID)

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!Ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "!Pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
