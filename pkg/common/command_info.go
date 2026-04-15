package common

import "github.com/bwmarrin/discordgo"

type CommandInfo struct {
	Name        string
	Description string
	Options     []*discordgo.ApplicationCommandOption
}
