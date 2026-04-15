package command

import (
	"fmt"
	"oat431/try-go-discord-bot/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func FactorialSlashInfo() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "factorial",
		Description: "Calculates the factorial of a number.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "number",
				Description: "The number to calculate factorial for.",
				Required:    true,
			},
		},
	}
}

func FactorialSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please provide a number to calculate factorial.",
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
	result := utils.Factorial(num)
	if result == -1 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Factorial is not defined for negative numbers: %d", num),
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("The factorial of %d is %d.", num, result),
			},
		})
	}
}
