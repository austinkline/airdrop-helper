package admin

import (
	"github.com/austinkline/airdrop/commands"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	refreshCommandName = "refresh"
	refreshCommand     = &discordgo.ApplicationCommand{
		Name:        refreshCommandName,
		Description: "Refresh all slash commands",
	}
)

func Refresh(s *discordgo.Session) string {
	response := ""
	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		response = "failed to get application commands"
		log.WithError(err).Error(response)
		return response
	}

	_, err = s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", cmds)
	if err != nil {
		response = "failed to overwrite application commands"
		log.WithError(err).Error(response)
		return response
	}

	return "Successfully refreshed commands"
}

func refreshCommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Info("refreshing commands")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Refreshing commands...",
		},
	})

	response := Refresh(s)
	log.Info("refreshed commands")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
}

func init() {
	commands.RegisterCommand(refreshCommand, refreshCommandName, refreshCommandsHandler)
}
