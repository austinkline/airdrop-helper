package handlers

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/austinkline/airdrop/network"
	"github.com/austinkline/airdrop/types"
	"github.com/bwmarrin/discordgo"
)

// Network is a Handler func
// It manages network configuration for various EVM endpoints
// that need to be accessed by the bot.
// It can add, remove, and list networks.
func Network(s *discordgo.Session, m *discordgo.Message, input string) (output string, err error) {
	params, err := parseFlags(input)
	if err != nil {
		return
	}

	if params.Get {
		if params.All {
			var networks []types.Network
			networks, err = network.GetAll()
			if err != nil {
				log.WithError(err).Error("failed to get networks")
				return
			} else {
				for _, n := range networks {
					output += n.Name + "\n"
				}
			}
		} else {
			var n types.Network
			n, err = network.Get(params.Name)
			if err != nil {
				log.WithError(err).Error("failed to get network")
				return
			} else {
				output = n.RpcURL
			}
		}
	} else if params.Add {
		n := types.Network{
			Name:   params.Name,
			RpcURL: params.RpcURL,
		}

		err = network.Add(n)
		if err != nil {
			log.WithError(err).Error("failed to add network")
		}

		var replyErr error
		if err == nil {
			output = "Network added"
		} else {
			output = "Failed to add network"
		}

		if replyErr != nil {
			log.WithError(replyErr).Error("failed to send reply")
		}

		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}

	return
}

type Params struct {
	Get    bool
	Add    bool
	All    bool
	Name   string
	RpcURL string
}

func parseFlags(input string) (p Params, err error) {
	flags := flag.NewFlagSet("network", flag.ContinueOnError)

	flags.BoolVar(&p.Get, "get", false, "whether to get the network")
	flags.BoolVar(&p.Add, "add", false, "whether to add a new network")
	flags.BoolVar(&p.All, "all", false, "whether to get all networks")
	flags.StringVar(&p.Name, "name", "", "the name of the network")
	flags.StringVar(&p.RpcURL, "url", "", "the RPC URL of the network")

	err = flags.Parse(strings.Split(input, " "))
	return
}

func init() {
	RegisterHandler("network", Network)
}
