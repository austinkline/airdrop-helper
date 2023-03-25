package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type Handler func(s *discordgo.Session, m *discordgo.Message, input string) (output string, err error)

// HandlerFuncs is a map of string to Handler funcs
// The string is the command that will trigger the handler
var HandlerFuncs = map[string]Handler{}

// StringToMap converts a single line string to a map
// it expects input like it might read from a command line
// and should return a string to string map with each field in the map
// corresponding to inputs in the original. Encasing inputs in quotes should
// allow for spaces in the input.
// Example:
// StringToMap("foo=bar baz=qux")
// returns map[string]string{
// 	"foo": "bar",
// 	"baz": "qux",
// }
func StringToMap(s string) map[string]string {
	m := map[string]string{}
	for _, v := range strings.Split(s, " ") {
		parts := strings.Split(v, "=")
		m[parts[0]] = parts[1]
	}
	return m
}

// MessageCreate
// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example, but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if the message does not start with !, ignore it
	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	// the first word of the message is the command, the rest is the input.
	// the first word should trim its leading !
	segments := strings.Split(m.Content, " ")
	command := segments[0][1:]
	input := strings.Join(segments[1:], " ")

	output, err := HandlerFuncs[command](s, m.Message, input)
	if err != nil {
		log.WithError(err).Error("error handling command")
	}

	s.ChannelMessageSend(m.ChannelID, output)
}

// RegisterHandler takes a string and a Handler func and adds it to the
// HandlerFuncs map. If the HandlerFuncs map does not exist, it will be initialized
// before adding the handler.
func RegisterHandler(command string, handler Handler) {
	if HandlerFuncs == nil {
		HandlerFuncs = map[string]Handler{}
	}
	HandlerFuncs[command] = handler
}
