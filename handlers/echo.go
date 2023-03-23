package handlers

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// EchoHandler is a func that is a Handler func
// It echos the contents of its message content back to the discordgo
// session that called it.
func EchoHandler(s *discordgo.Session, m *discordgo.Message, input string) error {
	// read the contents of input using a flag parser
	flags := flag.NewFlagSet("echo", flag.ContinueOnError)

	var (
		reverse = false
	)

	flags.BoolVar(&reverse, "reverse", false, "whether to reverse the input")
	err := flags.Parse(strings.Split(input, " "))
	if err != nil {
		return err
	}

	output := input
	// reverse the word order if the reverse flag is set
	if reverse {
		words := strings.Split(input, " ")
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		output = strings.Join(words, " ")
	}

	_, err = s.ChannelMessageSend(m.ChannelID, output)
	return err
}

func init() {
	RegisterHandler("echo", EchoHandler)
}
