package commands

import "github.com/bwmarrin/discordgo"

var (
	Commands        []*discordgo.ApplicationCommand
	CommandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
)

func RegisterCommand(command *discordgo.ApplicationCommand, name string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	// initialize Commands if it is nil
	if Commands == nil {
		Commands = make([]*discordgo.ApplicationCommand, 0)
	}

	// initialize CommandHandlers if it is nil
	if CommandHandlers == nil {
		CommandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
	}

	Commands = append(Commands, command)
	CommandHandlers[name] = handler
}
