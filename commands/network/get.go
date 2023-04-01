package network

import (
	"github.com/austinkline/airdrop/commands"
	"github.com/austinkline/airdrop/network"
	"github.com/bwmarrin/discordgo"
)

var (
	getNetworkCommandName = "getnetwork"
	getNetworkCommand     = &discordgo.ApplicationCommand{
		Name:        getNetworkCommandName,
		Description: "Get the url of a network",
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

func getNetworkHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	name := optionMap["name"].StringValue()

	var response string
	n, err := network.Get(name)
	if err == nil {
		response = n.RpcURL
	} else {
		response = "unable to get network"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
}

func init() {
	commands.RegisterCommand(getNetworkCommand, getNetworkCommandName, getNetworkHandler)
}
