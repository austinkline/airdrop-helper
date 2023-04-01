package account

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/austinkline/airdrop/address"
	"github.com/austinkline/airdrop/commands"
)

var (
	createAccountCommandName = "create-account"
	createAccountCommand     = &discordgo.ApplicationCommand{
		Name:        createAccountCommandName,
		Description: "Create a new account",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of the network to print",
				Required:    true,
			},
		},
	}
)

func createAccountHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	name := optionMap["name"].StringValue()

	addr, err := address.Create(name)
	responseMessage := "received"
	if err == nil {
		responseMessage = fmt.Sprintf("Created address %s with the name %s", addr, name)
	} else {
		log.WithError(err).Error("Failed to create address")
		responseMessage = "Failed to create address"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: responseMessage,
		},
	})
}

func init() {
	commands.RegisterCommand(createAccountCommand, createAccountCommandName, createAccountHandler)
}
