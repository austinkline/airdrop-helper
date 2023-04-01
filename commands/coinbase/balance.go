package coinbase

import (
	"fmt"
	"github.com/austinkline/airdrop/coinbase"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/austinkline/airdrop/commands"
)

var (
	cbBalanceCommandName = "cb-balance"
	createAccountCommand = &discordgo.ApplicationCommand{
		Name:        cbBalanceCommandName,
		Description: "Get the balance of a coinbase account",
	}
)

func cbBalanceHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	responseMessage := "received"

	account, err := coinbase.GetEthAccount()
	if err == nil {
		responseMessage = fmt.Sprintf("coinbase eth balance is %s", account.Balance.Amount)
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
