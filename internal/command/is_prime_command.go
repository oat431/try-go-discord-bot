package command

import (
	"fmt"
	"oat431/try-go-discord-bot/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func IsPrimeSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please provide a number to check.",
			},
		})
		return
	}

	numOption := options[0]
	if numOption.Type != discordgo.ApplicationCommandOptionInteger {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid input. Please provide an integer.",
			},
		})
		return
	}

	num := int(numOption.IntValue())
	if utils.IsPrime(num) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%d is a prime number.", num),
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%d is not a prime number.", num),
			},
		})
	}
}
