package handlers

import (
	"github.com/bwmarrin/discordgo"
	"sort"
	"strings"
)

// HelpHandler is a func that is a Handler func
// It prints a list of all available commands to the discordgo
// session that called it.
// The commands which can be called are the keys in HandlerFuncs
func HelpHandler(s *discordgo.Session, m *discordgo.Message, input string) (output string, err error) {
	// get all keys in the HandlerFuncs map
	keys := make([]string, 0, len(HandlerFuncs))
	for k := range HandlerFuncs {
		keys = append(keys, k)
	}

	// sort keys so that the output is consistent
	sort.Strings(keys)

	// the message to send back through the discord session consists of a
	// leading message saying "Available commands are:". Followed by each key in keys
	// on a new line.
	msg := "Available commands are:\n" + strings.Join(keys, "\n")

	//send msg as a message to the channel that m was sent from
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	return
}

func init() {
	RegisterHandler("help", HelpHandler)
}
