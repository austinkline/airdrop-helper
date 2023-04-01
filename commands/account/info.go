package account

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/austinkline/airdrop/address"
	"github.com/austinkline/airdrop/commands"
)

var (
	infoAccountCommandName = "get-account"
	infoAccountCommand     = &discordgo.ApplicationCommand{
		Name:        infoAccountCommandName,
		Description: "Get a created account's information",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of the account to print (must specify ",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "address",
				Description: "The address of the account to print",
				Required:    false,
			},
		},
	}
)

func infoAccountHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	nameOpt := optionMap["name"]
	addrOpt := optionMap["address"]

	var (
		name string;
		addr string
	)
	if nameOpt != nil {
		name = nameOpt.StringValue()
	}
	if addrOpt != nil {
		addr = addrOpt.StringValue()
	}

	responseMessage := "received"

	account := address.Account{}
	var err error

	if name != "" {
		account, err = address.GetAccountByName(name)
	} else if addr != "" {
		account, err = address.GetAccountByAddress(addr)
	} else {
		err = fmt.Errorf("must specify either name or address")
	}

	if err != nil {
		log.WithError(err).Error("Failed to get account")
		responseMessage = "Failed to get account"
	}

	if err == nil {
		responseMessage = fmt.Sprintf("Name: %s, Address: %s, PK: %s", account.Name, account.Address, account.PK)
	} else {
		log.WithError(err).Error("Failed to get account")
		responseMessage = "Failed to get account"
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
	commands.RegisterCommand(infoAccountCommand, infoAccountCommandName, infoAccountHandler)
}
