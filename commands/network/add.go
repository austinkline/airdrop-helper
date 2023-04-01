package network

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/austinkline/airdrop/commands"
	"github.com/austinkline/airdrop/network"
	"github.com/austinkline/airdrop/types"
)

var (
	addNetworkCommandName = "addnetwork"
	addNetworkCommand     = &discordgo.ApplicationCommand{
		Name:        addNetworkCommandName,
		Description: "Get the url of a network",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of the network to print",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "url",
				Description: "The rpc url of the network to use when interacting with it",
				Required:    true,
			},
		},
	}
)

func addNetworkHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	name := optionMap["name"].StringValue()
	url := optionMap["url"].StringValue()

	n := types.Network{
		Name:   name,
		RpcURL: url,
	}

	var responseMessage string
	err := network.Add(n)
	if err != nil {
		responseMessage = fmt.Sprintf("Failed to add network: %s", err)
	} else {
		responseMessage = fmt.Sprintf("Added network: %s", name)
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
	commands.RegisterCommand(addNetworkCommand, addNetworkCommandName, addNetworkHandler)
}
