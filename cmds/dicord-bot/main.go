package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/austinkline/airdrop/handlers"
)

const (
	envClientID = "DISCORD_CLIENT_ID"
	envToken    = "DISCORD_TOKEN"
)

var (
	discordToken string
)

func main() {
	discord, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(handlers.MessageCreate)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func init() {
	discordToken = os.Getenv(envToken)
}
