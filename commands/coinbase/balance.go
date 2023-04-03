package coinbase

import (
	"fmt"
	"github.com/austinkline/airdrop/coinbase"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/austinkline/airdrop/commands"
)

var (
	cbBalanceCommandName = "cb-balance"
	createAccountCommand = &discordgo.ApplicationCommand{
		Name:        cbBalanceCommandName,
		Description: "Get the balance of a coinbase account",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "symbol",
				Description: "The symbol of the coin to get the balance of",
				Required:    true,
			},
		},
	}
)

func cbBalanceHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Info("handling cb-balance...")
	responseMessage := "received"

	// get the symbol from command options
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	symbol := optionMap["symbol"].StringValue()
	symbol = strings.ToLower(symbol)

	id, err := coinbase.GetIdForSymbol(symbol)
	if err != nil {
		log.WithError(err).Error("Failed to get id for symbol")
		responseMessage = "No account found for symbol"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: responseMessage,
			},
		})
		return
	}

	account, err := coinbase.GetAccount(id)
	if err == nil {
		log.WithFields(log.Fields{"id": account.Id, "balance": account.Balance.Amount}).Info("got account")
		responseMessage = fmt.Sprintf("coinbase %s balance is %s", symbol, account.Balance.Amount)
	} else {
		log.WithError(err).Error("Failed to get balance")
		responseMessage = "Failed to get balance"
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
	commands.RegisterCommand(createAccountCommand, cbBalanceCommandName, cbBalanceHandler)
}
