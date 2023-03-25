package handlers

import (
	"database/sql"
	"flag"
	"github.com/bwmarrin/discordgo"
	"strings"

	database "github.com/austinkline/airdrop/db"
)

// StatusHandler is a func that is a Handler func
// It prints the status of the bot to the discordgo
// session that called it.
func StatusHandler(s *discordgo.Session, m *discordgo.Message, input string) (output string, err error) {
	// read the input using a flag parser

	var (
		all bool
		db  bool
	)

	flags := flag.NewFlagSet("status", flag.ContinueOnError)
	flags.BoolVar(&db, "db", false, "whether to check the database")
	flags.BoolVar(&all, "all", true, "whether to check the database")
	flags.Parse(strings.Split(input, " "))

	checks := []string{"I'm alive!"}

	if db || all {
		var c *sql.DB
		c, err = database.GetConnection()
		if err != nil {
			return
		}

		defer c.Close()

		err = c.Ping()
		if err != nil {
			checks = append(checks, "Database is down!")
		} else {
			checks = append(checks, "Database is up!")
		}
	}

	_, err = s.ChannelMessageSend(m.ChannelID, strings.Join(checks, "\n"))
	return
}

func init() {
	RegisterHandler("status", StatusHandler)
}
