package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/austinkline/airdrop/commands"
	"github.com/austinkline/airdrop/handlers"

	_ "github.com/austinkline/airdrop/commands/account"
	_ "github.com/austinkline/airdrop/commands/network"
)

const (
	envToken = "DISCORD_TOKEN"
)

// Bot parameters
var (
	BotToken = os.Getenv(envToken)

	GuildID        = *flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	RemoveCommands = *flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	ForceUpdate    = *flag.Bool("force", false, "Force update commands")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		registeredCommands[i] = cmd
	}

	s.AddHandler(handlers.HandleMessageReactAdd)

	if ForceUpdate {
		log.Println("Force updating commands...")
		for _, v := range registeredCommands {
			_, err := s.ApplicationCommandEdit(s.State.User.ID, GuildID, v.ID, v)
			if err != nil {
				log.Panicf("Cannot update '%v' command: %v", v.Name, err)
			}
		}
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if RemoveCommands {
		// Fetch all the bot's slash commands
		commands, err := s.ApplicationCommands(s.State.User.ID, GuildID)
		if err != nil {
			fmt.Println("Error fetching commands:", err)
			return
		}

		for _, c := range commands {
			err = s.ApplicationCommandDelete(s.State.User.ID, "", c.ID)
			if err != nil {
				fmt.Println("Error deleting commands:", err)
				return
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
