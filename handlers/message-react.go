package handlers

import (
	"fmt"
	"github.com/austinkline/airdrop/commands/admin"

	"github.com/bwmarrin/discordgo"
)

func HandleMessageReactAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}

	// Check if the reaction is the one we're interested in
	if r.Emoji.Name == "❌" {
		// Delete the message
		err := s.ChannelMessageDelete(r.ChannelID, r.MessageID)
		if err != nil {
			fmt.Println("Error deleting message:", err)
		}
	}

	// refresh all bot commands if the refresh emoji is used
	if r.Emoji.Name == "♻️" {
		response := admin.Refresh(s)
		_, _ = s.ChannelMessageSend(r.ChannelID, response)
	}
}
